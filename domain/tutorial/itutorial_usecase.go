package tutorial

import (
	"context"
	"gocrudssample/domain/tutorial/model"
)

type TutorialUsecaseInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error)
}
