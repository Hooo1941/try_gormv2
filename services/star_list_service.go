package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
)

type StarListService struct {
}

func (service *StarListService) List(user *models.User) *serializers.Response {
	var stars []models.Star
	if err := models.DB.
		Where("user_id=?", user.ID).
		Preload("User").Find(&stars).Error; err != nil {
		return &serializers.Response{
			Status:  e.SELECT_ERROR,
			Message: e.GetMsg(e.SELECT_ERROR),
			Error:   err.Error(),
		}
	}
	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data: serializers.BuildList(serializers.BuildStars(stars),
			uint(len(stars))),
	}
}
