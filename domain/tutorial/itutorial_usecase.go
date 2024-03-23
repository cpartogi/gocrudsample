package tutorial

import (
	"context"
	response "gocrudssample/schema/response"
)

type TutorialUsecaseInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret response.TutorialDetail, err error)
	GetTutorialTypes(ctx context.Context) (res []response.TutorialTypes, err error)
	GetTutorials(ctx context.Context, tutorialTypeId string) (res []response.TutorialList, err error)
}
