package main

import (
	"github.com/labstack/echo/v4"
	"hex/internal/adapters"
	"hex/internal/adapters/infra"
	"hex/internal/application/api"
	user "hex/internal/application/core"
)

func main() {
	userDbAdapter, err := infra.NewUserDbAdapter()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	user := user.New()
	apiUser := api.NewUserApplication(userDbAdapter, user)
	restUser := adapters.NewUserAdapter(apiUser)
	restUser.Run(e)

	err = e.Start(":8080")

	if err != nil {
		panic(err)
	}
}
