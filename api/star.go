package api

import (
	"comment/models"
	"comment/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateStar(c *gin.Context) {
	var service services.CreateStarService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, err.Error())
	} else {
		user := models.CurrentUser(c)
		res := service.Create(user)
		c.JSON(200, res)
	}
}

func DeleteStar(c *gin.Context) {
	var service services.DeleteStarService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, err.Error())
	} else {
		user := models.CurrentUser(c)
		res := service.Delete(user)
		c.JSON(200, res)
	}
}

func StarList(c *gin.Context) {
	var service services.StarListService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, err.Error())
	} else {
		user := models.CurrentUser(c)
		res := service.List(user)
		c.JSON(200, res)
	}
}

func Stared(c *gin.Context) {
	var service services.StaredService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, err.Error())
	} else {
		user := models.CurrentUser(c)
		articleID, _ := strconv.Atoi(c.Param("article_id"))
		res := service.Stared(user,uint(articleID))
		c.JSON(200, res)
	}
}
