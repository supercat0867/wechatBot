package wxhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// SendTextMsg 发送文本消息
func SendTextMsg(wxid, msg string) error {
	// 将数据转换为JSON格式
	jsonData, _ := json.Marshal(map[string]string{
		"wxid": wxid,
		"msg":  msg,
	})

	// 创建POST请求
	url := fmt.Sprintf("%s/SendTextMsg", os.Getenv("WECHAT_DOMAIN"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 解析响应
	var responseJSON SendTextMsgResponse
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		return fmt.Errorf("响应解析失败：%v\n", err)
	}

	if responseJSON.SendTextMsg != "1" {
		return fmt.Errorf("消息发送失败")
	}

	return nil
}
