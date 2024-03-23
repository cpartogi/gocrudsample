package main

import (
	"gocrudsample/domain/tutorial/repo"
	appInit "gocrudsample/init"
	"net/http"
	"time"

	handler "gocrudsample/domain/tutorial/handler/http"
	usecase "gocrudsample/domain/tutorial/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	rdb := appInit.NewRedisClient()

	defer rdb.Close()

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
	tutorialRepo := repo.NewTutorialRepo(dbRead.DB, dbWrite, rdb)

	tutorialUc := usecase.NewTutorialUsecase(tutorialRepo, timeoutContext)

	// End of DI Stepss

	handler.NewTutorialHandler(e, tutorialUc)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
