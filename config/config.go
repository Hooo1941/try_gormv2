/*
 * @Date: 2020-11-21 23:26:09
 * @LastEditors: QiuJhao
 * @LastEditTime: 2020-11-29 00:41:26
 */
package config

import (
	"github.com/gin-gonic/gin"
	//"comment/cache"
	"comment/models"
	"comment/util"
	"os"

	"github.com/joho/godotenv"
)

var CurrentTime string

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	CurrentTime = "2006-01-02 15:04:05"
	if err := godotenv.Load(); err != nil {
		util.Log().Panic("环境变量文件加载失败", err)
	}
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	models.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()
}
