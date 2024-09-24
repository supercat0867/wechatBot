package glm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var AccessTokenInfo struct {
	AccessToken string
	Exp         time.Time
}

// GenerateAccessToken 生成AccessToken
func GenerateAccessToken() error {
	apiKey := os.Getenv("GLM_API_KEY")
	parts := strings.Split(apiKey, ".")
	if len(parts) != 2 {
		return errors.New("KEY格式错误")
	}

	id := parts[0]
	secret := parts[1]

	// 定义 1 小时的过期时间
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建 JWT 的 payload
	claims := jwt.MapClaims{
		"api_key":   id,
		"exp":       expirationTime.Unix(),
		"timestamp": time.Now().UnixMilli(),
	}

	// 创建 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 设置头信息
	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"

	// 使用 secret 签名生成 token
	accessToken, err := token.SignedString([]byte(secret))
	AccessTokenInfo.AccessToken = accessToken
	AccessTokenInfo.Exp = expirationTime
	return err
}

// GetAccessToken 获取AccessToken
func GetAccessToken() string {
	// 检查token是否过期
	if AccessTokenInfo.AccessToken == "" || time.Until(AccessTokenInfo.Exp) < time.Minute {
		// 过期了，重新生成
		if err := GenerateAccessToken(); err != nil {
			log.Printf("GLM AccessToken生成失败：%s", err)
		}
	}
	return AccessTokenInfo.AccessToken
}

// CallGLM 调用GLM模型
// accessToken: 访问令牌
// model: 模型名称，例如：glm-4-plus、glm-4-0520、glm-4、glm-4-air
// messages: 对话消息列表
func CallGLM(model string, messages []Message) (string, error) {
	messages = append([]Message{{Content: SystemPrompt, Role: "system"}}, messages...)

	// 将数据转换为JSON格式
	jsonData, _ := json.Marshal(map[string]interface{}{
		"model":    model,
		"messages": messages,
	})

	// 创建POST请求
	url := "https://open.bigmodel.cn/api/paas/v4/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GetAccessToken())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析响应
	var responseJSON CallGLMResponse
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		return "", fmt.Errorf("响应解析失败：%v\n", err)
	}

	return responseJSON.Choices[0].Message.Content, nil
}

// CallCogView 调用CogView模型
// prompt: 提示词
func CallCogView(prompt string) (string, error) {
	// 将数据转换为JSON格式
	jsonData, _ := json.Marshal(map[string]interface{}{
		"model":  "cogview-3-plus",
		"prompt": prompt,
	})

	// 创建POST请求
	url := "https://open.bigmodel.cn/api/paas/v4/images/generations"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GetAccessToken())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析响应
	var responseJSON CallCogViewResponse
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		return "", fmt.Errorf("响应解析失败：%v\n", err)
	}

	return responseJSON.Data[0].Url, nil
}
