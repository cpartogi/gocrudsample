package usecase

import (
	"context"
	"gocrudssample/domain/tutorial"
	response "gocrudssample/schema/response"
	"time"

	"gocrudssample/lib/constant"

	"github.com/google/uuid"
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

func (u *TutorialUsecase) GetDetailTutorial(ctx context.Context, tutorialId string) (res response.TutorialDetail, err error) {

	_, err = uuid.Parse(tutorialId)
	if err != nil {
		return res, constant.ErrInvalidUuid
	}

	ret, err := u.tutorialRepo.GetDetailTutorial(ctx, tutorialId)

	if err != nil {
		return res, err
	}

	var lastUpdate time.Time
	if ret.UpdatedAt != nil {
		lastUpdate = *ret.UpdatedAt
	} else {
		lastUpdate = ret.CreatedAt
	}

	res = response.TutorialDetail{
		Id:           ret.Id,
		Title:        ret.Title,
		TutorialType: ret.TutorialTypeName,
		Keywords:     ret.Keywords,
		Sequence:     ret.Sequence,
		Description:  ret.Description,
		LastUpdate:   lastUpdate.Format(time.RFC3339),
	}

	return res, nil
}
