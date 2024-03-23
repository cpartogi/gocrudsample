package tutorial

import (
	"context"
	"gocrudssample/domain/tutorial/model"
)

type TutorialRepoInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error)
	GetTutorialTypes(ctx context.Context) (ret []model.TutorialTypes, err error)
	GetTutorials(ctx context.Context, tutorialTypeId string) (ret []model.Tutorials, err error)
	AddTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
}
