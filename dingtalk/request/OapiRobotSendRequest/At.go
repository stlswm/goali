package OapiRobotSendRequest

type At struct {
	AtMobiles []string
	IsAtAll   bool
}

// 设置唯一的提示手机号
func (a *At) SetAtMobile(mobile string) *At {
	a.AtMobiles = []string{mobile}
	return a
}

// 设置提示手机号
func (a *At) SetAtMobiles(mobiles []string) *At {
	a.AtMobiles = mobiles
	return a
}

// 追加提示手机号
func (a *At) AppendAtMobile(mobile string) *At {
	a.AtMobiles = append(a.AtMobiles, mobile)
	return a
}

// 提示所有人
func (a *At) AtAll() *At {
	a.AtMobiles = []string{}
	a.IsAtAll = true
	return a
}
