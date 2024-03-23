package tutorial

import (
	"context"
	response "gocrudssample/schema/response"
)

type TutorialUsecaseInterface interface {
	GetDetailTutorial(ctx context.Context, tutorialId string) (ret response.TutorialDetail, err error)
}
