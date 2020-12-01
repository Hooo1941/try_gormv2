package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
)

type DeleteStarService struct {
	ArticleID uint `json:"article_id" form:"article_id"`
}

func (service *DeleteStarService) Delete(user *models.User) *serializers.Response {
	if service.ArticleID == 0 {
		return &serializers.Response{
			Status:  e.INPUT_EMPTY,
			Message: e.GetMsg(e.INPUT_EMPTY),
		}
	}
	star := models.Star{
		UserID:    user.ID,
		ArticleID: service.ArticleID,
	}
	var stars []models.Star
	var total int64
	if err := models.DB.Model(stars).
		Where("article_id=? and user_id=?", service.ArticleID, user.ID).
		Count(&total).Error; err != nil {
		return &serializers.Response{
			Status:  e.SELECT_ERROR,
			Message: e.GetMsg(e.SELECT_ERROR),
			Error:   err.Error(),
		}
	}
	if total == 0 {
		return &serializers.Response{
			Status:  e.NO_STAR,
			Message: e.GetMsg(e.NO_STAR),
		}
	}
	if err := models.DB.Delete(&star).Error; err != nil {
		return &serializers.Response{
			Status:  e.DELETE_ERROR,
			Message: e.GetMsg(e.DELETE_ERROR),
			Error:   err.Error(),
		}
	}
	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
	}
}
