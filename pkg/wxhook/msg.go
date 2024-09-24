package wxhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// OpenMsgReceiver 开启实时消息接受功能
func OpenMsgReceiver(url string) error {
	// 将数据转换为JSON格式
	jsonData, _ := json.Marshal(map[string]string{
		"isEnable": "1",
		"url":      url,
	})

	// 创建POST请求
	endPoint := fmt.Sprintf("%s/ConfigureMsgRecive", os.Getenv("WECHAT_DOMAIN"))
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonData))
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
	var responseJSON ConfigureMsgReciveResponse
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		return fmt.Errorf("响应解析失败：%v\n", err)
	}

	if responseJSON.ConfigureMsgRecive != "1" {
		return fmt.Errorf("开启消息接受功能失败")
	}

	return nil
}

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

// SendImgMsg 发送图片消息
func SendImgMsg(wxid, picPath string) error {
	// 将数据转换为JSON格式
	jsonData, _ := json.Marshal(map[string]string{
		"wxid":        wxid,
		"picpath":     picPath,
		"diyfilename": "1.jpg",
	})

	// 创建POST请求
	url := fmt.Sprintf("%s/SendPicMsg", os.Getenv("WECHAT_DOMAIN"))
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
	var responseJSON SendPicMsgResponse
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		return fmt.Errorf("响应解析失败：%v\n", err)
	}

	if responseJSON.SendPicMsg != "1" {
		return fmt.Errorf("消息发送失败")
	}

	return nil
}
