package main

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/repository"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfgPtr, err := config.NewConfig()
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		fmt.Println(errMessage)
		return
	}

	e, err := InitApp(cfgPtr)
	port := fmt.Sprintf(":%d", cfgPtr.Api.Port)
	e.Logger.Fatal(e.Start(port))
}

func InitApp(cfg *config.Config) (*echo.Echo, error) {
	e := echo.New()

	db, err := repository.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	repo := repository.NewUserRepository(db)
	h := api.NewUserHandler(repo)
	api.RegisterHandlers(e, h)

	return e, nil
}
