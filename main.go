package main

import (
	"log"
	"net/http"
	"os"
	"time"

	controller "github.com/Friske2/pokemon-api/controller/v1"
	repo "github.com/Friske2/pokemon-api/repository"
	service "github.com/Friske2/pokemon-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	r := gin.Default()
	dsn := "host=localhost user=sa password=P@ssw0rd dbname=Pokemon port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	fileInfo, err := openLogFile("./log/" + getFilenameDate())
	if err != nil {
		log.Fatal(err)
	}
	newLogger := logger.New(
		log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

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

func getFilenameDate() string {
	// Use layout string for time format.
	const layout = "01-02-2006"
	// Place now in the string.
	t := time.Now()
	return "file-" + t.Format(layout) + ".log"
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
