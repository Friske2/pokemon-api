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

	// r.GET("/pokemon", func(c *gin.Context) {
	// 	monters := []domain.Monters{}
	// 	// db.Raw("SELECT * FROM monters").Scan(&monters)
	// 	db.Where("enabled = ?", true).Find(&monters)
	// 	c.JSON(http.StatusOK, monters)
	// })

	// r.POST("/pokemon", func(c *gin.Context) {
	// 	body := dto.MonterValue{}
	// 	err := c.ShouldBind(&body)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "bad request",
	// 		})
	// 		return
	// 	}
	// 	monters := domain.Monters{
	// 		Name:        body.Name,
	// 		CreatedDate: time.Now(),
	// 		CreatedBy:   "system",
	// 		Enabled:     true,
	// 	}
	// 	res := db.Create(&monters)
	// 	if res.Error != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": res.Error.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":   "success",
	// 		"monsterId": monters.ID,
	// 	})
	// })

	// r.PUT("/pokemon/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	body := dto.MonterValue{}
	// 	err := c.ShouldBind(&body)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	monterId, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	monters := domain.Monters{}
	// 	m := db.First(&monters, monterId)
	// 	if m.Error != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"message": m.Error.Error(),
	// 		})
	// 		return
	// 	}
	// 	var updateBody = map[string]interface{}{
	// 		"Name": body.Name,
	// 	}
	// 	if body.Enabled != nil {
	// 		updateBody["Enabled"] = *body.Enabled
	// 	}
	// 	res := db.Model(&monters).Where("id = ?", monterId).Updates(updateBody)

	// 	if res.Error != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": res.Error.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":   "updated success",
	// 		"monsterId": monters.ID,
	// 	})
	// })

	// r.GET("/pokemon/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	monters := domain.Monters{}
	// 	monterId, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	res := db.First(&monters, monterId)
	// 	if res.Error != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"message": res.Error.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, monters)
	// })

	// r.DELETE("/pokemon/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	monters := domain.Monters{}
	// 	monterId, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	res := db.First(&monters, monterId)
	// 	if res.Error != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"message": res.Error.Error(),
	// 		})
	// 		return
	// 	}
	// 	res = db.Delete(domain.Monters{}, monterId)
	// 	if res.Error != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": res.Error.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "deleted success",
	// 	})
	// })
}
