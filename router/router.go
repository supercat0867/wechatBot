package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"wechatBot/domain/msg"
	"wechatBot/pkg/wxhook"
)

// RegisterMessageHandler 注册消息处理器
func RegisterMessageHandler(r *gin.Engine) {
	// 开启消息处理
	err := wxhook.OpenMsgReceiver(fmt.Sprintf("%s/wechat/msgHandler", os.Getenv("SERVER_DOMAIN")))
	if err != nil {
		panic(err)
	}
	// 处理消息
	r.POST("/wechat/msgHandler", msg.MessageHandler)
}
