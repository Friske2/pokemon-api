package v1

import (
	"net/http"
	"strconv"

	"github.com/Friske2/pokemon-api/domain"
	"github.com/Friske2/pokemon-api/dto"
	"github.com/Friske2/pokemon-api/model"
	"github.com/gin-gonic/gin"
)

type montersController struct {
	MonterService domain.MontersService
}

func NewMontersController(c *gin.Engine, m domain.MontersService) {
	controller := &montersController{
		MonterService: m,
	}

	v1 := c.Group("/api/v1")
	{
		v1.GET("/pokemon", controller.FindAll)
		v1.GET("/pokemon/:id", controller.GetById)
		v1.POST("/pokemon", controller.Insert)
		v1.PUT("/pokemon/:id", controller.Update)
		v1.DELETE("/pokemon/:id", controller.Delete)
	}
}

func (c *montersController) FindAll(ctx *gin.Context) {
	// find all
	monters, err := c.MonterService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, monters)
}

func (c *montersController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	monterId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	monter, err := c.MonterService.GetById(int32(monterId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, monter)
}

func (c *montersController) Insert(ctx *gin.Context) {
	// insert
	body := dto.MonterValue{}
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	monters := model.Monter{
		Name: body.Name,
	}
	id, err := c.MonterService.Insert(monters)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"monsterId": id,
	})
}

func (c *montersController) Update(ctx *gin.Context) {
	// update
	id := ctx.Param("id")
	body := dto.MonterValue{}
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	monterId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	_, err = c.MonterService.GetById(int32(monterId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = c.MonterService.Update(int32(monterId), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "updated success",
		"monsterId": monterId,
	})

}

func (c *montersController) Delete(ctx *gin.Context) {
	// delete
	id := ctx.Param("id")
	monterId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	_, err = c.MonterService.GetById(int32(monterId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = c.MonterService.Delete(int32(monterId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted success",
	})
}
