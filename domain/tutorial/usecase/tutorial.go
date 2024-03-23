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

func (u *TutorialUsecase) GetTutorialTypes(ctx context.Context) (res []response.TutorialTypes, err error) {

	ret, err := u.tutorialRepo.GetTutorialTypes(ctx)

	if err != nil {
		return res, err
	}

	for _, rt := range ret {
		res = append(res, response.TutorialTypes{
			Id:       rt.Id,
			TypeName: rt.TypeName,
		})
	}

	return
}

func (u *TutorialUsecase) GetTutorials(ctx context.Context, tutorialTypeId string) (res []response.TutorialList, err error) {

	if tutorialTypeId != "" {
		_, err = uuid.Parse(tutorialTypeId)
		if err != nil {
			return res, constant.ErrInvalidUuid
		}
	}

	ret, err := u.tutorialRepo.GetTutorials(ctx, tutorialTypeId)

	if err != nil {
		return res, err
	}

	for _, rt := range ret {
		res = append(res, response.TutorialList{
			Id:           rt.Id,
			Title:        rt.Title,
			TutorialType: rt.TutorialTypeName,
		})
	}

	return
}
