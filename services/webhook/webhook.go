package webhook

import (
	"Topicgram/services/bot"

	"github.com/gin-gonic/gin"
)

func NewEngine() (router *gin.Engine) {
	router = gin.New()
	router.SetTrustedProxies([]string{"0.0.0.0/0", "::/0"})

	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())

	// 使用 bot 包提供的路径获取函数，确保路径已规范化并可自定义
	router.POST(string(bot.WebHookPath()), bot.HookHandler)
	return
}
