package tutorial

import (
	"context"
	"gocrudsample/domain/tutorial/model"
	response "gocrudsample/schema/response"
)

type TutorialUsecaseInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret response.TutorialDetail, err error)
	GetTutorialTypes(ctx context.Context) (res []response.TutorialTypes, err error)
	GetTutorials(ctx context.Context, tutorialTypeId string) (res []response.TutorialList, err error)
	AddTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
	UpdateTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
	DeleteTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
	PatchTutorial(ctx context.Context, tutorial model.Tutorials) (err error)
}
