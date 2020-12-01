package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
	"github.com/gin-gonic/gin"
)

type StaredService struct {
}

func (service *StaredService) Stared(user *models.User, ArticleID uint) *serializers.Response {
	if ArticleID == 0 {
		return &serializers.Response{
			Status:  e.INPUT_EMPTY,
			Message: e.GetMsg(e.INPUT_EMPTY),
		}
	}
	var stars []models.Star
	var total int64
	if err := models.DB.Model(stars).
		Where("article_id=? and user_id=?", ArticleID, user.ID).
		Count(&total).Error; err != nil {
		return &serializers.Response{
			Status:  e.SELECT_ERROR,
			Message: e.GetMsg(e.SELECT_ERROR),
			Error:   err.Error(),
		}
	}
	if total == 0 {
		return &serializers.Response{
			Status:  e.SUCCESS,
			Message: e.GetMsg(e.SUCCESS),
			Data: gin.H{
				"star": 0,
			},
		}
	}
	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data: gin.H{
			"star": 1,
		},
	}
}
