package msg

// ReceiveMsg 消息接收格式
type ReceiveMsg struct {
	ServerPort string        `json:"ServerPort"` // 服务端口
	SelfWxID   string        `json:"selfwxid"`   // 个人ID
	SendOrrecv string        `json:"sendorrecv"` // 收发类型
	MsgNumber  string        `json:"msgnumber"`  // 消息数量
	MsgList    []MessageList `json:"msglist"`    // 消息列表
}
type MessageList struct {
	MsgType   string `json:"msgtype"`   // 消息类型
	MsgSvrid  string `json:"msgsvrid"`  // 服务器消息ID
	Msg       string `json:"msg"`       // 消息内容
	FromType  string `json:"fromtype"`  // 个人消息=1 群消息=2
	FromID    string `json:"fromid"`    // 发送方微信ID
	FromName  string `json:"fromname"`  // 发送方微信名
	FromGID   string `json:"fromgid"`   // 发送方群聊ID
	FromGName string `json:"fromgname"` // 发送方群聊名称
	ToID      string `json:"toid"`      // 接收方微信ID
	ToName    string `json:"toname"`    // 接收方微信名
}

type ActionInput struct {
	Prompt string `json:"prompt"`
}
