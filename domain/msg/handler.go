package msg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"wechatBot/pkg/glm"
	"wechatBot/pkg/wxhook"
)

func MessageHandler(g *gin.Context) {
	var req ReceiveMsg
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
	}

	// 只处理收到的消息
	if req.SendOrrecv == "2" {
		for _, msg := range req.MsgList {
			fmt.Println("===========================================================================================")
			fmt.Printf("消息类型：%s\n", msg.MsgType)
			fmt.Printf("消息内容：%s\n", msg.Msg)
			fmt.Printf("消息来源：%s\n", msg.FromType)
			fmt.Printf("发送方wxid：%s\n", msg.FromID)
			fmt.Printf("发送方昵称：%s\n", msg.FromName)
			fmt.Printf("发送方群聊id：%s\n", msg.FromGID)
			fmt.Printf("发送方群聊名称：%s\n", msg.FormGName)
			fmt.Printf("接收方wxid:%s\n", msg.ToID)
			fmt.Printf("接收方昵称：%s\n", msg.ToName)
			fmt.Println("===========================================================================================")

			if msg.FromType == "1" {
				// 好友消息
				if msg.MsgType == "1" {
					// 处理文本消息
					reply, err := glm.CallGLM("glm-4-plus", []glm.Message{{Content: msg.Msg, Role: "user"}})
					if err != nil {
						log.Printf("GLM 调用失败：%v", err)
						reply = "GLM模型调用失败！"
					}
					_ = wxhook.SendTextMsg(msg.FromID, reply)
				} else {
					_ = wxhook.SendTextMsg(msg.FromID, "暂不支持此消息类型")
				}
			} else if msg.FromType == "2" {
				// 群聊消息
				if msg.MsgType == "1" && strings.Contains(msg.Msg, "@旺柴") {
					// 只监听@自己的消息的文本消息
					content := strings.ReplaceAll(msg.Msg, "@旺柴", "")
					content = strings.TrimSpace(content)
					// 处理文本消息
					reply, err := glm.CallGLM("glm-4-plus", []glm.Message{{Content: msg.Msg, Role: "user"}})
					if err != nil {
						log.Printf("GLM 调用失败：%v", err)
						reply = "GLM模型调用失败！"
					}
					_ = wxhook.SendTextMsg(msg.FromGID, reply)
				}
			}
		}

	}

	g.JSON(200, gin.H{"code": 0, "msg": "success"})
	return
}
