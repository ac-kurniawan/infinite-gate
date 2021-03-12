package postgres

import (
	"fmt"
	"infinitegate/src/modules/postgres/models"
	"infinitegate/src/util/config"
	"infinitegate/src/util/debug"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Config config.PostgresConnection
}

func (p *Connection) setup() (*gorm.DB, error) {
	address := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d" +
		" sslmode=disable", p.Config.Host, p.Config.Username, p.Config.Password,
		p.Config.Name, p.Config.Port)
	db, err := gorm.Open(postgres.Open(address), &gorm.Config{})
	if err != nil {
		debug.Error("[Database] - [Postgres]", err.Error())
		return nil, err
	}

	dbs, _ := db.DB()

	dbs.SetMaxOpenConns(10)
	dbs.SetMaxIdleConns(10)
	dbs.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func (p *Connection) Connect() *gorm.DB {
	connection, _ := p.setup()

	connection.SetupJoinTable(&models.User{}, "Projects", &models.UsersProject{})
	connection.AutoMigrate(&models.User{}, &models.Project{}, &models.UsersProject{})

	debug.Info("[Database] - [Postgres]", "Connected!")

	return connection
}

func NewPGConnection(pgConnection config.PostgresConnection) Connection {
	return Connection{pgConnection}
}

