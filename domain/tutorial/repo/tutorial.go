package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"gocrudssample/domain/tutorial"
	"gocrudssample/domain/tutorial/model"
	"gocrudssample/lib/constant"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

type TutorialRepo struct {
	db   *sql.DB
	gopg *pg.DB
	rdb  *redis.Client
}

func NewTutorialRepo(db *sql.DB, gopg *pg.DB, rdb *redis.Client) tutorial.TutorialRepoInterface {
	return &TutorialRepo{
		db:   db,
		gopg: gopg,
		rdb:  rdb,
	}
}

func (r *TutorialRepo) GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error) {

	//check redis
	tutorialKey := `tutorial_%s`

	key := fmt.Sprintf(tutorialKey, tutorialId)

	var res string
	cr := r.rdb.Get(key)
	if err = cr.Err(); err != nil {

		if err == redis.Nil {

			query := `SELECT t.id, t.tutorial_type_id, t.keywords, t.sequence, t.title, t.description, ttype.type_name, t.created_at, t.updated_at FROM tutorials t, tutorial_types ttype WHERE t.tutorial_type_id = ttype.id AND t.id = $1 AND t.deleted_at is null`

			row := r.db.QueryRowContext(ctx, query, tutorialId)

			err = row.Scan(&ret.Id, &ret.TutorialTypeId, &ret.Keywords, &ret.Sequence, &ret.Title, &ret.Description, &ret.TutorialTypes.TypeName, &ret.CreatedAt, &ret.UpdatedAt)

			if err != nil {

				if err == sql.ErrNoRows {
					return ret, constant.ErrNotFound
				} else {
					return
				}
			}

			expDuration := time.Hour * 1
			b, err := json.Marshal(ret)
			if err != nil {
				return ret, err
			}

			err = r.rdb.Set(key, b, expDuration).Err()
			if err != nil {
				return ret, err
			}

		}
	} else {

		res, err = cr.Result()
		if err != nil {
			return
		}

		err = json.Unmarshal([]byte(res), &ret)
		if err != nil {
			return ret, err
		}
	}
	return
}

func (r *TutorialRepo) GetTutorialTypes(ctx context.Context) (ret []model.TutorialTypes, err error) {
	//check redis
	key := `tutorialTypes`

	var res string
	cr := r.rdb.Get(key)
	if err = cr.Err(); err != nil {

		if err == redis.Nil {

			query := `SELECT id, type_name FROM tutorial_types WHERE deleted_at is null`

			rows, err := r.db.QueryContext(ctx, query)

			if err != nil {
				return ret, err
			}

			defer rows.Close()

			for rows.Next() {
				var resp model.TutorialTypes
				err = rows.Scan(&resp.Id, &resp.TypeName)

				if err != nil {
					return ret, err
				}

				ret = append(ret, resp)

			}

			expDuration := time.Hour * 1
			b, err := json.Marshal(ret)
			if err != nil {
				return ret, err
			}

			err = r.rdb.Set(key, b, expDuration).Err()
			if err != nil {
				return ret, err
			}

		}
	} else {

		res, err = cr.Result()
		if err != nil {
			return
		}

		err = json.Unmarshal([]byte(res), &ret)
		if err != nil {
			return ret, err
		}
	}
	return
}

func (r *TutorialRepo) GetTutorials(ctx context.Context, tutorialTypeId string) (ret []model.Tutorials, err error) {

	var query string
	var rows *sql.Rows

	if tutorialTypeId != "" {
		query = `SELECT t.id, t.title, ttype.type_name FROM tutorials t, tutorial_types ttype WHERE t.tutorial_type_id = ttype.id AND t.tutorial_type_id = $1 AND t.deleted_at is null ORDER BY t.sequence, ttype.type_name `

		rows, err = r.db.QueryContext(ctx, query, tutorialTypeId)
	} else {
		query = `SELECT t.id, t.title, ttype.type_name FROM tutorials t, tutorial_types ttype WHERE t.tutorial_type_id = ttype.id AND t.deleted_at is null ORDER BY t.sequence `

		rows, err = r.db.QueryContext(ctx, query)
	}

	if err != nil {
		return ret, err
	}

	defer rows.Close()

	for rows.Next() {
		var resp model.Tutorials
		err = rows.Scan(&resp.Id, &resp.Title, &resp.TutorialTypes.TypeName)

		if err != nil {
			return ret, err
		}

		ret = append(ret, resp)

	}

	return
}

func (r *TutorialRepo) AddTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	tx, err := r.gopg.Begin()
	if err != nil {
		return
	}

	_, err = tx.ModelContext(ctx, &tutorial).Insert()
	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r *TutorialRepo) UpdateTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	tx, err := r.gopg.Begin()
	if err != nil {
		return
	}

	_, err = tx.ModelContext(ctx, &tutorial).Column("tutorial_type_id", "keywords", "sequence", "title", "description", "updated_by", "updated_at").WherePK().Update()
	if err != nil {
		tx.Rollback()
		return
	}

	// delete redis
	tutorialKey := `tutorial_%s`

	key := fmt.Sprintf(tutorialKey, tutorial.Id)

	_, err = r.rdb.Del(key).Result()
	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	return

}

func (r *TutorialRepo) DeleteTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	tx, err := r.gopg.Begin()
	if err != nil {
		return
	}

	_, err = tx.ModelContext(ctx, &tutorial).Column("deleted_by", "deleted_at").WherePK().Update()
	if err != nil {
		tx.Rollback()
		return
	}

	// delete redis
	tutorialKey := `tutorial_%s`

	key := fmt.Sprintf(tutorialKey, tutorial.Id)

	_, err = r.rdb.Del(key).Result()
	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	return

}
