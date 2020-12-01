package models

import "time"

const CommentEveryPageCount int = 500

type Comment struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	User      User       `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	UserID    uint       `gorm:"not null" sql:"type:uint REFERENCES users(id)"`
	ReplyTo   User       `gorm:"foreignkey:ReplyToID;association_foreignkey:ID"`
	ReplyToID uint       `sql:"default:null" sql:"type:uint REFERENCES users(id)"`
	Article   article    `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
	ArticleID uint       `gorm:"not null" sql:"type:uint REFERENCES articles(id)"`
	Parent    *Comment   `gorm:"foreignkey:ParentID;association_foreignkey:ID;"`
	ParentID  uint       `sql:"default:null" sql:"type:uint REFERENCES comments(id)"`
	Root      *Comment   `gorm:"foreignkey:RootID;association_foreignkey:ID;"`
	RootID    uint       `sql:"default:null" sql:"type:uint REFERENCES comments(id)"`
	Content   string     `gorm:"not null"`
	Replys    []Comment  `sql:"default:null"`
	LikeCount uint       `gorm:"default:0"`
	IsLike    bool       `gorm:"default:false"`
}

type Star struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	User      User    `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	UserID    uint    `gorm:"not null" sql:"type:uint REFERENCES users(id)"`
	Article   article `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
	ArticleID uint    `gorm:"not null" sql:"type:uint REFERENCES articles(id)"`
}
