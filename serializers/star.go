package serializers

import (
	"comment/config"
	"comment/models"
)

type StarSerializer struct {
	ArticleID uint   `json:"article_id"`
	CreatedAt string `json:"created_at"`
}

func BuildStar(item models.Star) StarSerializer {
	return StarSerializer{
		ArticleID: item.ArticleID,
		CreatedAt: item.CreatedAt.Format(config.CurrentTime),
	}
}

func BuildStars(items []models.Star) []StarSerializer {
	var stars []StarSerializer
	for _, item := range items {
		stars = append(stars, BuildStar(item))
	}
	return stars
}
