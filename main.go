package main

import (
	"gocrudssample/domain/tutorial/repo"
	appInit "gocrudssample/init"
	"net/http"
	"time"

	handler "gocrudssample/domain/tutorial/handler/http"
	usecase "gocrudssample/domain/tutorial/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

func init() {
	// Start pre-requisite app dependencies
	appInit.StartAppInit()
}

func main() {

	dbRead, err := appInit.ConnectToPGServerRead()
	if err != nil {
		log.S().Fatal(err)
	}

	defer dbRead.DB.Close()

	dbWrite, err := appInit.ConnectToPGServerWrite()
	if err != nil {
		log.S().Fatal(err)
	}

	defer dbWrite.Close()

	// init router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy")
	})

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	//DI: Repository & Usecase
	tutorialRepo := repo.NewTutorialRepo(dbRead.DB, dbWrite)

	tutorialUc := usecase.NewTutorialUsecase(tutorialRepo, timeoutContext)

	// End of DI Stepss

	handler.NewTutorialHandler(e, tutorialUc)

	//e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
