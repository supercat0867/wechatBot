package msg

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"regexp"
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
			fmt.Printf("发送方群聊名称：%s\n", msg.FromGName)
			fmt.Printf("接收方wxid:%s\n", msg.ToID)
			fmt.Printf("接收方昵称：%s\n", msg.ToName)
			fmt.Println("===========================================================================================")

			if msg.FromType == "1" {
				// 好友消息
				if msg.MsgType == "1" {
					// 处理文本消息
					DealTextMsg(msg.Msg, msg.FromID)
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
					DealTextMsg(msg.Msg, msg.FromGID)
				}
			}
		}

	}

	g.JSON(200, gin.H{"code": 0, "msg": "success"})
	return
}

// 提取Action和ActionInput
func extractActionAndParams(input string) (string, string, error) {
	re := regexp.MustCompile(`(?s)Action:\s*(.*?)\s*Action Input:\s*({.*})`)
	// 执行正则表达式匹配
	matches := re.FindStringSubmatch(input)

	// 检查匹配结果
	if len(matches) > 2 {
		action := matches[1]
		actionInput := matches[2]
		return action, actionInput, nil
	} else {
		return "", "", fmt.Errorf("匹配失败")
	}
}

// DealTextMsg 处理文本消息
func DealTextMsg(msg, fromId string) {
	reply, err := glm.CallGLM("glm-4-plus", []glm.Message{{Content: msg, Role: "user"}})
	if err != nil {
		log.Printf("GLM 调用失败：%v", err)
		reply = "GLM模型调用失败！"
	}

	// 提取参数
	action, actionInput, err := extractActionAndParams(reply)
	if err == nil {
		// 解析actionInput
		var args ActionInput
		if err = json.Unmarshal([]byte(actionInput), &args); err != nil {
			_ = wxhook.SendTextMsg(fromId, fmt.Sprintf("actionInput 解析失败：%v", err))
		} else {
			// 根据响应来调用不同的工具
			if action == "文生图" {
				// 调用文生图工具
				src, err := glm.CallCogView(args.Prompt)
				if err != nil {
					_ = wxhook.SendTextMsg(fromId, fmt.Sprintf("文生图调用失败：%v", err))
				} else {
					// 发送图片消息
					_ = wxhook.SendImgMsg(fromId, src, fmt.Sprintf("%s.jpg", uuid.New().String()))
				}
			} else {
				_ = wxhook.SendTextMsg(fromId, fmt.Sprintf("暂不支持此工具【%s】", action))
			}
		}
	} else {
		// 不存在工具，直接回复文本消息
		_ = wxhook.SendTextMsg(fromId, reply)
	}
}
