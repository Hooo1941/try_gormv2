package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
	"strings"
)

type RegisterService struct {
	UserName        string `json:"username" form:"username"`
	PassWord        string `json:"password" form:"password"`
	PassWordConfirm string `json:"passwordConfirm" form:"passwordConfirm"`
}

func (service *RegisterService) Valid() *serializers.Response {
	if service.UserName == "" || service.PassWord == "" || service.PassWordConfirm == "" {
		return &serializers.Response{
			Status:  e.INPUT_EMPTY,
			Message: e.GetMsg(e.INPUT_EMPTY),
		}
	}
	if strings.Count(service.UserName, "")-1 < 3 {
		return &serializers.Response{
			Status:  e.USERNAME_TOO_SHORT,
			Message: e.GetMsg(e.USERNAME_TOO_SHORT),
		}
	}
	if strings.Count(service.UserName, "")-1 > 20 {
		return &serializers.Response{
			Status:  e.USERNAME_TOO_LONG,
			Message: e.GetMsg(e.USERNAME_TOO_LONG),
		}
	}
	if strings.Count(service.PassWord, "")-1 < 6 {
		return &serializers.Response{
			Status:  e.PASSWORD_TOO_SHORT,
			Message: e.GetMsg(e.PASSWORD_TOO_SHORT),
		}
	}
	if strings.Count(service.PassWord, "")-1 > 20 {
		return &serializers.Response{
			Status:  e.PASSWORD_TOO_LONG,
			Message: e.GetMsg(e.PASSWORD_TOO_LONG),
		}
	}

	if service.PassWord != service.PassWordConfirm {
		return &serializers.Response{
			Status:  e.USER_PASSWORD_NOT_CONSISTENT,
			Message: e.GetMsg(e.USER_PASSWORD_NOT_CONSISTENT),
		}
	}

	var count int64
	models.DB.Model(&models.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializers.Response{
			Status:  e.USERNAME_ALREADY_EXISTED,
			Message: e.GetMsg(e.USERNAME_ALREADY_EXISTED),
		}
	}
	return nil
}

func (service *RegisterService) Register() *serializers.Response {
	user := models.User{
		UserName: service.UserName,
	}
	// valid check
	if response := service.Valid(); response != nil {
		return response
	}
	// set user password
	if err := user.SetPassWord(service.PassWord); err != nil {
		return &serializers.Response{
			Status:  e.USER_SET_PASSWORD_ERROR,
			Message: e.GetMsg(e.USER_SET_PASSWORD_ERROR),
		}
	}
	// create user
	if err := models.DB.Create(&user).Error; err != nil {
		return &serializers.Response{
			Status:  e.REGISTER_ERROR,
			Message: e.GetMsg(e.REGISTER_ERROR),
		}
	}
	// register ok
	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data:    serializers.BuildUser(user),
	}
}
