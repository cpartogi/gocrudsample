package http

import (
	"gocrudsample/domain/tutorial"
	"gocrudsample/domain/tutorial/model"

	"gocrudsample/lib/constant"

	"gocrudsample/lib/pkg/utils"

	"gocrudsample/schema/request"

	"github.com/labstack/echo/v4"
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

	e.POST("/tutorials", handler.Addtutorial)
	e.PUT("/tutorials/:tutorial_id", handler.UpdateTutorial)
	e.DELETE("/tutorials/:tutorial_id", handler.DeleteTutorial)
	e.PATCH("/tutorials/:tutorial_id", handler.PatchTutorial)
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

func (h *TutorialHandler) Addtutorial(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.AddTutorial

	c.Bind(&req)

	tutorialsData := model.Tutorials{
		TutorialTypeId: req.TutorialTypeId,
		Keywords:       req.Keywords,
		Sequence:       req.Sequence,
		Title:          req.Title,
		Description:    req.Description,
	}

	err := h.tutorialUsecase.AddTutorial(ctx, tutorialsData)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessAddData, "")

}

func (h *TutorialHandler) UpdateTutorial(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.AddTutorial
	c.Bind(&req)

	tutorialId := c.Param("tutorial_id")

	tutorialsData := model.Tutorials{
		Id:             tutorialId,
		TutorialTypeId: req.TutorialTypeId,
		Keywords:       req.Keywords,
		Sequence:       req.Sequence,
		Title:          req.Title,
		Description:    req.Description,
	}

	err := h.tutorialUsecase.UpdateTutorial(ctx, tutorialsData)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessUpdateData, "")

}

func (h *TutorialHandler) DeleteTutorial(c echo.Context) error {
	ctx := c.Request().Context()

	tutorialId := c.Param("tutorial_id")

	tutorialsData := model.Tutorials{
		Id: tutorialId,
	}

	err := h.tutorialUsecase.DeleteTutorial(ctx, tutorialsData)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessDeleteData, "")

}

func (h *TutorialHandler) PatchTutorial(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.PatchTutorial
	c.Bind(&req)

	tutorialId := c.Param("tutorial_id")

	tutorialsData := model.Tutorials{
		Id:    tutorialId,
		Title: req.Title,
	}

	err := h.tutorialUsecase.PatchTutorial(ctx, tutorialsData)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessUpdateData, "")

}
