package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"gocrudssample/domain/tutorial"
	"gocrudssample/domain/tutorial/model"

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

			err = row.Scan(&ret.Id, &ret.TutorialTypeId, &ret.Keywords, &ret.Sequence, &ret.Title, &ret.Description, &ret.TutorialTypeName, &ret.CreatedAt, &ret.UpdatedAt)

			if err != nil {
				return
			}

			expDuration := time.Hour * 2
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
	return ret, nil
}
