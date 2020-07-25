# dingtalkrus

DingTalk Hook for [Logrus](https://github.com/Sirupsen/logrus).

### Use

```go
package main

import (
	"github.com/leafney/dingtalkrus"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stderr)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.AddHook(dingtalkrus.NewHook(
		"", // dingtalk token
		"", // dingtalk secret
		dingtalkrus.LevelThreshold(logrus.ErrorLevel)),
	)

	logrus.Info("This is the info test message.")
	logrus.WithFields(dingtalkrus.SendTextMsg("This is the warn test message.",[]string{},false)).Warn()
	logrus.WithFields(dingtalkrus.SendMarkdownMsg("杭州天气","#### 杭州天气 \n 9度，西北风1级，空气良89，相对温度73%\n",[]string{},false)).Error()
}
```

### Installation

```
go get github.com/leafney/dingtalkrus
```

### Message parameters

SendTextMsg

| 参数 | 参数类型 | 必须 | 说明 |
| --- | ------ | ---- | --- |
| content | String | 是 | 消息内容 |
| atMobiles | Array | 否 | 被@人的手机号(在content里添加@人的手机号) |
| isAtAll | Boolean | 否 | 是否@所有人 |

SendMarkdownMsg

| 参数 | 参数类型 | 必须 | 说明 |
| --- | ------ | ---- | --- |
| title | String | 是 | 首屏会话透出的展示内容 |
| text | String | 是 | markdown格式的消息 |
| atMobiles | Array | 否 | 被@人的手机号（在text内容里需要有@手机号） |
| isAtAll | Boolean | 否 | 是否@所有人 |

SendLinkMsg

| 参数 | 参数类型 | 必须 | 说明 |
| --- | ------ | ---- | --- |
| title | String | 是 | 消息标题 |
| text | String 是 | 消息内容。如果太长只会部分展示 |
| messageUrl | String | 是 | 点击消息跳转的URL |
| picUrl | String | 否 | 图片URL |

### Reference

* [dingrus](https://github.com/dandans-dan/dingrus)
* [slackrus](https://github.com/johntdyer/slackrus)