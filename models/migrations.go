package models

import "comment/util"

//执行数据迁移

func migration() {
	// 自动迁移模式
	err := DB.AutoMigrate(&Article{}, &User{}, &Comment{}, &Star{})
	if err != nil {
		util.Log().Panic("创建数据库不成功", err)
	}
	//DB.Model(&Comment{}).AddForeignKey("user_id", "users(id)",
	//	"CASCADE", "CASCADE")
	//DB.Model(&Comment{}).AddForeignKey("article_id", "articles(id)",
	//	"CASCADE", "CASCADE")
	//DB.Model(&Comment{}).AddForeignKey("parent_id", "comments(id)",
	//	"CASCADE", "CASCADE")
	//DB.Model(&Comment{}).AddForeignKey("root_id", "comments(id)",
	//	"CASCADE", "CASCADE")
	//DB.Model(&Comment{}).AddForeignKey("reply_to_id", "users(id)",
	//	"CASCADE", "CASCADE")
	//DB.Model(&Star{}).AddForeignKey("user_id", "users(id)",
	//	"CASCADE", "CASCADE")
	//DB.Model(&Star{}).AddForeignKey("article_id", "articles(id)",
	//	"CASCADE", "CASCADE")
}
