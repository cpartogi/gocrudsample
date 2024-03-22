package usecase

import (
	"context"
	"gocrudssample/domain/tutorial"
	"gocrudssample/domain/tutorial/model"
	"time"
)

type TutorialUsecase struct {
	tutorialRepo   tutorial.TutorialRepoInterface
	contextTimeout time.Duration
}

func NewTutorialUsecase(tutorialRepo tutorial.TutorialRepoInterface, timeout time.Duration) tutorial.TutorialUsecaseInterface {
	return &TutorialUsecase{
		tutorialRepo:   tutorialRepo,
		contextTimeout: timeout,
	}
}

func (u *TutorialUsecase) GetDetailTutorial(ctx context.Context, tutorialId string) (ret model.Tutorials, err error) {

	ret, err = u.tutorialRepo.GetDetailTutorial(ctx, tutorialId)

	if err != nil {
		return ret, err
	}

	return ret, nil
}
