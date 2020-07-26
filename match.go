package dingtalkrus

import "regexp"

func dingTalkMsgFilter(b []byte) bool {
	m,err:= regexp.Match(`msgtype\":\"(text)|(markdown)|(link)\"`,b)
	if err!=nil{
		return false
	}
	return m
}