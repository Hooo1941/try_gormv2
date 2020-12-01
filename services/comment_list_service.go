package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
	"strconv"
)

type CommentListService struct {
}

func (service *CommentListService) List(articleID uint, p string) *serializers.Response {
	if articleID == 0 {
		return &serializers.Response{
			Status:  e.INPUT_EMPTY,
			Message: e.GetMsg(e.INPUT_EMPTY),
		}
	}
	page, _ := strconv.Atoi(p)
	start := models.CommentEveryPageCount * (page - 1)

	// 获取顶级评论并做分页
	var comments []models.Comment
	var Replys [][]models.Comment
	if err := models.DB.Limit(models.CommentEveryPageCount).Offset(start).
		Where("article_id=? and parent_id is null", articleID).
		Preload("User").Find(&comments).Error; err != nil {
		return &serializers.Response{
			Status:  e.SELECT_ERROR,
			Message: e.GetMsg(e.SELECT_ERROR),
			Error:   err.Error(),
		}
	}

	//	获取顶级评论数量
	var total int64
	if err := models.DB.Model(comments).
		Where("article_id=? and parent_id is null", articleID).
		Count(&total).Error; err != nil {
		return &serializers.Response{
			Status:  e.SELECT_ERROR,
			Message: e.GetMsg(e.SELECT_ERROR),
			Error:   err.Error(),
		}
	}

	//　获取当前评论下方的所有回复，当前评论的点赞情况
	for i := 0; i < len(comments); i++ {
		var rep []models.Comment
		if err := models.DB.Where("root_id=?", comments[i].ID).
			Preload("User").Preload("ReplyTo").
			Find(&rep).Error; err != nil {
			return &serializers.Response{
				Status:  e.SELECT_ERROR,
				Message: e.GetMsg(e.SELECT_ERROR),
				Error:   err.Error(),
			}
		}
		Replys = append(Replys, rep)
	}

	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data: serializers.BuildList(serializers.BuildComments(comments,Replys),
			uint(total)),
	}
}
