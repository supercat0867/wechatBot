package glm

import (
	"fmt"
	"regexp"
	"testing"
	"wechatBot/pkg/wxhook"
)

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

func TestCallGLM(t *testing.T) {
	reply, err := CallGLM("glm-4-plus", []Message{{Content: "帮我画一只猫咪", Role: "user"}})
	if err != nil {
		t.Errorf("GLM 调用失败：%v", err)
	}
	t.Logf("GLM 回复：%s", reply)
	action, actionInput, err := extractActionAndParams(reply)
	if err != nil {
		t.Errorf("提取Action和ActionInput失败：%v", err)
	} else {
		t.Logf("Action: %s", action)
		t.Logf("Action Input: %s", actionInput)
	}
}

func TestCallCogView(t *testing.T) {
	reply, err := CallCogView("一只可爱的猫咪，毛色柔软，眼神明亮，正舒服地蜷缩在温暖的阳光下。它的耳朵微微向前倾斜，似乎在倾听周围的声音。背景是一片温馨的室内环境，有柔软的地毯和舒适的沙发，整体色调温暖而柔和，细节丰富，画风偏向于写实主义。")
	if err != nil {
		t.Errorf("CogView 调用失败：%v", err)
	} else {
		t.Logf("CogView 回复：%s", reply)
		// 发送图片
		if err := wxhook.SendImgMsg("wxid_ry29s202j5q912", reply); err != nil {
			t.Errorf("发送图片失败：%v", err)
		}
	}
}
