package main

import (
	"net/http"

	controller "github.com/Friske2/pokemon-api/controller/v1"
	repo "github.com/Friske2/pokemon-api/repository"
	service "github.com/Friske2/pokemon-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	dsn := "host=localhost user=sa password=P@ssw0rd dbname=Pokemon port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	montersRepo := repo.NewMontersRepo(db)
	monterService := service.NewMontersService(montersRepo)
	controller.NewMontersController(r, monterService)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
