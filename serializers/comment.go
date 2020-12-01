package serializers

import (
	"comment/config"
	"comment/models"
)

type CommentSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	ArticleID uint              `json:"article_id"`
	Content   string            `json:"content"`
	CreatedAt string            `json:"created_at"`
	Replys    []ReplySerializer `json:"replys"`
	LikeCount uint              `json:"like_count"`
	IsLike    bool              `json:"is_like"`
}

type ReplySerializer struct {
	ID        uint           `json:"id"`
	User      UserSerializer `json:"user"`
	ReplyTo   UserSerializer `json:"reply_to"`
	ParentID  uint           `json:"parent_id"`
	RootID    uint           `json:"root_id"`
	Content   string         `json:"content"`
	CreatedAt string         `json:"created_at"`
	LikeCount uint              `json:"like_count"`
	IsLike    bool              `json:"is_like"`
}

func BuildReply(item models.Comment) ReplySerializer {
	return ReplySerializer{
		ID:        item.ID,
		User:      BuildUser(item.User),
		ReplyTo:   BuildUser(item.ReplyTo),
		ParentID:  item.ParentID,
		RootID:    item.RootID,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Format(config.CurrentTime),
	}
}

func BuildComment(item models.Comment, reply []models.Comment) CommentSerializer {
	var replys []ReplySerializer
	if len(reply) != 0 {
		for _, reply := range reply {
			replys = append(replys, BuildReply(reply))
		}
	}
	return CommentSerializer{
		ID:        item.ID,
		User:      BuildUser(item.User),
		ArticleID: item.ArticleID,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Format(config.CurrentTime),
		Replys:    replys,
		LikeCount: item.LikeCount,
		IsLike:    item.IsLike,
	}
}

func BuildComments(items []models.Comment, replys [][]models.Comment) []CommentSerializer {
	var comments []CommentSerializer
	for i, item := range items {
		comments = append(comments, BuildComment(item, replys[i]))
	}
	return comments
}
