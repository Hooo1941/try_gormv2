package models

import "time"

const CommentEveryPageCount int = 500

type Comment struct {
	ID        uint `gorm:"primary_key autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	User      User
	UserID    uint       `gorm:"not null"`
	ReplyTo   User
	ReplyToID uint       `gorm:"default:null"`
	Article   Article
	ArticleID uint       `gorm:"not null"`
	Parent    *Comment
	ParentID  uint       `gorm:"default:null"`
	Root      *Comment
	RootID    uint       `gorm:"default:null"`
	Content   string     `gorm:"not null"`
	LikeCount uint       `gorm:"default:0"`
	IsLike    bool       `gorm:"default:false"`
}

type Star struct {
	ID        uint `gorm:"primary_key autoIncrement"`
	CreatedAt time.Time
	User      User
	UserID    uint    `gorm:"not null"`
	Article   Article
	ArticleID uint    `gorm:"not null"`
}
