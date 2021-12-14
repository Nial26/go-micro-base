package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go-micro-base/config"
	"go-micro-base/item"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// server will have shared dependencies for the whole server
type server struct {
	db     *gorm.DB
	router *gin.Engine
}

func createDatabase(dbc config.DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
							dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to database : %v", err)
		os.Exit(1)
	}
	db.AutoMigrate(&item.Entity{})
	return db
}

func (s *server) Start() {
	s.registerRoutes()
	s.router.Run()
}


func New(config config.Config) server {
	return server{
		db:     createDatabase(config.DbConfig),
		router: gin.Default(),
	}
}




