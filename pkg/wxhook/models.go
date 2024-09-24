package wxhook

// ConfigureMsgReciveResponse 配置消息接收响应
type ConfigureMsgReciveResponse struct {
	ConfigureMsgRecive string `json:"ConfigureMsgRecive"`
}

// SendTextMsgResponse 发送文本消息响应
type SendTextMsgResponse struct {
	SendTextMsg string `json:"SendTextMsg"`
}

// SendPicMsgResponse 发送图片消息响应
type SendPicMsgResponse struct {
	SendPicMsg string `json:"SendPicMsg"`
}
