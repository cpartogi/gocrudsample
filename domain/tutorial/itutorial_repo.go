package tutorial

import (
	"context"
	"gocrudssample/domain/tutorial/model"
)

type TutorialRepoInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error)
}
