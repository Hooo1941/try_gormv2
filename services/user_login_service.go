package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
	"comment/util"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func (service *LoginService) Login() *serializers.Response {
	if service.UserName == "" || service.PassWord == "" {
		return &serializers.Response{
			Status:  e.INPUT_EMPTY,
			Message: e.GetMsg(e.INPUT_EMPTY),
		}
	}
	var user models.User
	if err := models.DB.Where("user_name=?", service.UserName).Find(&user).Error; err != nil {
		return &serializers.Response{
			Status:  e.USERNAME_OR_PASSWORD_ERROR,
			Message: e.GetMsg(e.USERNAME_OR_PASSWORD_ERROR),
		}
	}
	if user.CheckPassword(service.PassWord) == false {
		return &serializers.Response{
			Status:  e.USERNAME_OR_PASSWORD_ERROR,
			Message: e.GetMsg(e.USERNAME_OR_PASSWORD_ERROR),
		}
	}

	token := util.GenerateToken(user.ID, user.UserName)
	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data: gin.H{
			"userID": user.ID,
			"token":  token,
		},
	}
}
