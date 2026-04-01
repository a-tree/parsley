package main

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/repository"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		fmt.Println(errMessage)
		return
	}

	e := echo.New()
	db, err := repository.NewDB(cfg)
	if err != nil {
		fmt.Printf("[main] repository.NewDB Error %v", err)
		return
	}
	repo := repository.NewUserRepository(db)
	h := api.NewUserHandler(repo)
	api.RegisterHandlers(e, h)

	port := fmt.Sprintf(":%d", cfg.Api.Port)
	e.Logger.Fatal(e.Start(port))
}
