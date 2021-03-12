package auth

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"infinitegate/src/services/auth"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	authImpl auth.IAuthService
}

func (c *Controller) CreateToken(ctx echo.Context) error {

	// Get body data
	bodyBytes, _ := ioutil.ReadAll(ctx.Request().Body)
	var body map[string]interface{}
	json.Unmarshal(bodyBytes, &body)

	email := body["email"].(string)
	name := body["name"].(string)
	userId, _ := body["userId"].(float64)

	token, err := c.authImpl.CreateToken(name, email, int(userId))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "501",
			"payload": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"code":    "200",
		"payload": token,
	})
}

func NewController(impl auth.IAuthService) *Controller {
	return &Controller{
		authImpl: impl,
	}
}
