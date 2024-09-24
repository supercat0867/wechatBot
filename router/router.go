package router

import (
	"github.com/gin-gonic/gin"
	"wechatBot/domain/msg"
)

// RegisterMessageHandler 注册消息处理器
func RegisterMessageHandler(r *gin.Engine) {
	r.POST("/wechat/msgHandler", msg.MessageHandler)
}
