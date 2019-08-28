package OapiRobotSendRequest

// 链接消息
type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
}

func (l *Link) SetTitle(title string) *Link {
	l.Title = title
	return l
}
func (l *Link) SetText(text string) *Link {
	l.Text = text
	return l
}
func (l *Link) SetMessageUrl(messageUrl string) *Link {
	l.MessageUrl = messageUrl
	return l
}
func (l *Link) SetPicUrl(picUrl string) *Link {
	l.PicUrl = picUrl
	return l
}
