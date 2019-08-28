package OapiRobotSendRequest

// 文本消息
type Text struct {
	At
	Content string
}

// 设置文本内容
func (t *Text) SetContent(content string) *Text {
	t.Content = content
	return t
}
