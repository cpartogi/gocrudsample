package repo

import (
	"context"
	"database/sql"

	"gocrudssample/domain/tutorial"
	"gocrudssample/domain/tutorial/model"

	"github.com/go-pg/pg"
)

type TutorialRepo struct {
	db   *sql.DB
	gopg *pg.DB
}

func NewTutorialRepo(db *sql.DB, gopg *pg.DB) tutorial.TutorialRepoInterface {
	return &TutorialRepo{
		db:   db,
		gopg: gopg,
	}
}

func (r *TutorialRepo) GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error) {

	query := `SELECT t.id, t.tutorial_type_id, t.keywords, t.sequence, t.title, t.description, ttype.type_name FROM tutorials t, tutorial_types ttype WHERE t.tutorial_type_id = ttype.id AND t.id = $1 AND t.deleted_at is null`

	row := r.db.QueryRowContext(ctx, query, tutorialId)

	err = row.Scan(&ret.Id, &ret.TutorialTypeId, &ret.Keywords, &ret.Sequence, &ret.Title, &ret.Description, &ret.TutorialTypeName)

	if err != nil {
		return
	}

	return ret, nil
}
