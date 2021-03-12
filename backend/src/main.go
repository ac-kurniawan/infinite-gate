package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"infinitegate/src/app/http"
	authModule "infinitegate/src/modules/api/auth"
	projectModule "infinitegate/src/modules/api/project"
	"infinitegate/src/modules/postgres"
	"infinitegate/src/modules/postgres/repositories"
	"infinitegate/src/services/auth"
	"infinitegate/src/services/project"
	"infinitegate/src/util/config"
	"infinitegate/src/util/debug"
)

func newProjectService(db *gorm.DB) project.IProjectService {
	projectRepository := repositories.NewProjectPG(db)
	projectService := project.NewProjectRepositoryImpl(&projectRepository)

	return projectService
}

func newAuthService() auth.IAuthService {
	return auth.NewAuthService("rahasia")
}

func main() {
	fmt.Printf("API Service Newgate\n")

	cfg := config.Config{}
	cfg.ReadFromDotEnv()

	// Connect to DB
	pg := postgres.NewPGConnection(cfg.PGConfig)
	db := pg.Connect()
	dbCtx, _ := db.DB()

	defer func(dbCtx *sql.DB) {
		err := dbCtx.Close()
		if err != nil {
			debug.Error("[Database] - [Postgres]", err.Error())
		}
	}(dbCtx)

	// Create service
	projectService := newProjectService(db)
	authService := newAuthService()

	// Create controller
	projectController := projectModule.NewController(projectService)
	authController := authModule.NewController(authService)

	// Create router
	router := http.ControllerMap{
		Project: projectController,
		Auth:    authController,
	}

	// Starting HTTP server
	address := fmt.Sprintf("0.0.0.0:8030")
	fmt.Printf("Start HTTP server on %s\n", address)
	http.HttpServer(address, router)

}
