package project

import (
	"github.com/dgrijalva/jwt-go"
	"infinitegate/src/services/project"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	projectImlp project.IProjectService
}

func (c *Controller) InsertProjectByUserID(ctx echo.Context) error {
	IP := new(InsertProjectRequest)
	if err := ctx.Bind(&IP); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			ErrorResponseProject{4000103, err.Error()},
		)
	}

	if err := ctx.Validate(IP); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			ErrorResponseProject{4000103, err.Error()},
		)
	}

	// Get userID from JWT payload
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["userId"].(float64)

	projectObj := project.Project{
		Name:        IP.Name,
		AccessLevel: IP.AccessLevel,
	}
	result, err := c.projectImlp.InsertProjectByUserID(int(id), projectObj)

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			ErrorResponseProject{4000103, err.Error()},
		)
	}

	res := DataResponse{
		ID:          result.ID,
		Name:        result.Name,
		AccessLevel: result.AccessLevel,
		CreatedAt:   result.CreatedAt,
	}

	return ctx.JSON(http.StatusOK, SuccessResponseFind{200, res})
}

func (c *Controller) FindProjectByID(ctx echo.Context) error {
	// Get userID from JWT payload
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["userId"].(float64)

	idInt, _ := strconv.Atoi(ctx.Param("id"))

	result, err := c.projectImlp.FindProjectByID(int(id), idInt)

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			ErrorResponseProject{4000101, err.Error()},
		)
	}

	if result.Name == "" {
		return ctx.JSON(
			http.StatusBadRequest,
			ErrorResponseProject{4000101, "project not found"},
		)
	}

	res := DataResponse{
		ID:          result.ID,
		Name:        result.Name,
		AccessLevel: result.AccessLevel,
		CreatedAt:   result.CreatedAt,
	}

	return ctx.JSON(http.StatusOK, SuccessResponseFind{200, res})
}

func (c *Controller) FindProjectsByUserID(ctx echo.Context) error {
	// Get userID from JWT payload
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["userId"].(float64)

	// find projects by userId
	projects, err := c.projectImlp.FindProjectsByUserID(int(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			ErrorResponseProject{4000102, err.Error()},
		)
	}

	var responseMap []DataResponse

	for _, v := range projects {
		responseMap = append(
			responseMap, DataResponse{
				ID:          v.ID,
				Name:        v.Name,
				AccessLevel: v.AccessLevel,
				CreatedAt:   v.CreatedAt,
			},
		)
	}

	return ctx.JSON(http.StatusOK, SuccessResponseFinds{200, responseMap})
}

func NewController(impl project.IProjectService) *Controller {
	return &Controller{
		projectImlp: impl,
	}
}
