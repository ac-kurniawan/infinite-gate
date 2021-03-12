package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"infinitegate/src/modules/api/auth"
	"infinitegate/src/modules/api/project"
	"infinitegate/src/util/debug"
	"net/http"
)

type ControllerMap struct {
	Project *project.Controller
	Auth    *auth.Controller
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func HttpServer(address string, controllerMap ControllerMap) {
	e := echo.New()

	// Register validator module
	e.Validator = &CustomValidator{validator: validator.New()}

	// Auth endpoint
	authEndpoint := e.Group("v1/auth")
	authEndpoint.POST("", controllerMap.Auth.CreateToken)

	// Project endpoint
	projectEndpoint := e.Group("v1/project")
	projectEndpoint.Use(middleware.JWT([]byte("rahasia")))
	projectEndpoint.GET("/:id", controllerMap.Project.FindProjectByID)
	projectEndpoint.GET("", controllerMap.Project.FindProjectsByUserID)
	projectEndpoint.POST("", controllerMap.Project.InsertProjectByUserID)

	// health check endpoint
	e.GET("/health", func(context echo.Context) error {
		return context.String(http.StatusOK, "OK")
	})

	// Start to listening
	if err := e.Start(address); err != nil {
		debug.Error("[HTTP]", err.Error())
	}
}
