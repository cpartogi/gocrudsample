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

	e.GET("/tutorials/:tutorial_id", handler.GetDetailTutorial)
	e.GET("/tutorials/types", handler.GetTutorialTypes)
	e.GET("/tutorials", handler.GetTutorials)
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

func (h *TutorialHandler) GetTutorialTypes(c echo.Context) error {

	ctx := c.Request().Context()

	res, err := h.tutorialUsecase.GetTutorialTypes(ctx)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, res)

}

func (h *TutorialHandler) GetTutorials(c echo.Context) error {

	ctx := c.Request().Context()
	tutorialTypeId := c.QueryParam("tutorialTypeId")

	res, err := h.tutorialUsecase.GetTutorials(ctx, tutorialTypeId)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, res)

}
