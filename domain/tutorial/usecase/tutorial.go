package usecase

import (
	"context"
	"gocrudssample/domain/tutorial"
	"gocrudssample/domain/tutorial/model"
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
		TutorialType: ret.TutorialTypes.TypeName,
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
			TutorialType: rt.TutorialTypes.TypeName,
		})
	}

	return
}

func (u *TutorialUsecase) AddTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	// validation
	_, err = uuid.Parse(tutorial.TutorialTypeId)
	if err != nil {
		return constant.ErrInvalidUuid
	}

	tutorialTypes, _ := u.tutorialRepo.GetTutorialTypes(ctx)

	mapTutorialType := map[string]*model.TutorialTypes{}

	for _, ttype := range tutorialTypes {
		mapTutorialType[ttype.Id] = &model.TutorialTypes{
			Id:       ttype.Id,
			TypeName: ttype.TypeName,
		}
	}

	if mapTutorialType[tutorial.TutorialTypeId] == nil {
		return constant.ErrTypeNotFound
	}

	if tutorial.Title == "" {
		return constant.ErrTitle
	}

	tutorial.Id = uuid.New().String()
	tutorial.CreatedAt = time.Now().UTC()
	tutorial.CreatedBy = "Admin"

	err = u.tutorialRepo.AddTutorial(ctx, tutorial)

	if err != nil {
		return
	}

	return
}

func (u *TutorialUsecase) UpdateTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	// validation
	_, err = uuid.Parse(tutorial.Id)
	if err != nil {
		return constant.ErrInvalidUuid
	}

	_, err = uuid.Parse(tutorial.TutorialTypeId)
	if err != nil {
		return constant.ErrInvalidUuid
	}

	tutorialTypes, _ := u.tutorialRepo.GetTutorialTypes(ctx)

	mapTutorialType := map[string]*model.TutorialTypes{}

	for _, ttype := range tutorialTypes {
		mapTutorialType[ttype.Id] = &model.TutorialTypes{
			Id:       ttype.Id,
			TypeName: ttype.TypeName,
		}
	}

	if mapTutorialType[tutorial.TutorialTypeId] == nil {
		return constant.ErrTypeNotFound
	}

	if tutorial.Title == "" {
		return constant.ErrTitle
	}

	updatedAt := time.Now().UTC()
	updatedBy := "Admin"

	tutorial.UpdatedAt = &updatedAt
	tutorial.UpdatedBy = &updatedBy

	err = u.tutorialRepo.UpdateTutorial(ctx, tutorial)

	if err != nil {
		return
	}

	return
}

func (u *TutorialUsecase) DeleteTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	// validation
	_, err = uuid.Parse(tutorial.Id)
	if err != nil {
		return constant.ErrInvalidUuid
	}

	deletedAt := time.Now().UTC()
	deletedBy := "Admin"

	tutorial.DeletedAt = &deletedAt
	tutorial.DeletedBy = &deletedBy

	err = u.tutorialRepo.DeleteTutorial(ctx, tutorial)

	if err != nil {
		return
	}

	return
}

func (u *TutorialUsecase) PatchTutorial(ctx context.Context, tutorial model.Tutorials) (err error) {

	// validation
	_, err = uuid.Parse(tutorial.Id)
	if err != nil {
		return constant.ErrInvalidUuid
	}

	updatedAt := time.Now().UTC()
	updatedBy := "Admin"

	tutorial.UpdatedAt = &updatedAt
	tutorial.UpdatedBy = &updatedBy

	err = u.tutorialRepo.PatchTutorial(ctx, tutorial)

	if err != nil {
		return
	}

	return
}
