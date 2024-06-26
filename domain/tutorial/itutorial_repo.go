package tutorial

import (
	"context"
	"gocrudsample/domain/tutorial/model"
)

type TutorialRepoInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error)
	GetTutorialTypes(ctx context.Context) (ret []model.TutorialTypes, err error)
	GetTutorials(ctx context.Context, tutorialTypeId string) (ret []model.Tutorials, err error)
	AddTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
	UpdateTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
	DeleteTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
	PatchTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
}
