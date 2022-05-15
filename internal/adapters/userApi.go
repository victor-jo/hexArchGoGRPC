package adapters

import (
	"github.com/labstack/echo/v4"
	"hex/internal/ports"
	"net/http"
	"strconv"
)

type UserAdapter struct {
	api ports.UserApiPorts
}

func NewUserAdapter(api ports.UserApiPorts) *UserAdapter {
	return &UserAdapter{
		api: api,
	}
}

func (adpater UserAdapter) Run(e *echo.Echo) {
	e.GET("/user/:id", func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err == nil {
			return c.String(http.StatusOK, "can't not parameter parsed.")
		}

		user, err := adpater.api.Find(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
	})

	e.POST("/user", func(c echo.Context) error {
		params := make(map[string]string)
		_ = c.Bind(&params)

		username, password := params["username"], params["password"]
		if len(username) == 0 || len(password) == 0 {
			return c.String(http.StatusOK, "can't not parameter parsed.")
		}

		id, err := adpater.api.Join(username, password)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, id)
	})
}
