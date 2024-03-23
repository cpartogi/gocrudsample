package http

import (
	"gocrudssample/domain/tutorial"

	"gocrudssample/lib/constant"

	"gocrudssample/lib/pkg/utils"

	"github.com/labstack/echo"
)

type TutorialHandler struct {
	tutorialUsecase tutorial.TutorialUsecaseInterface
}

func NewTutorialHandler(e *echo.Echo, us tutorial.TutorialUsecaseInterface) {
	handler := &TutorialHandler{
		tutorialUsecase: us,
	}

	e.GET("/tutorial/:tutorial_id", handler.GetDetailTutorial)
}

func (h *TutorialHandler) GetDetailTutorial(c echo.Context) error {

	ctx := c.Request().Context()
	tutorialId := c.Param("tutorial_id")

	res, err := h.tutorialUsecase.GetDetailTutorial(ctx, tutorialId)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, res)

}
