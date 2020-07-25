package dingtalkrus

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestDingTalkHook(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stderr)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.AddHook(NewHook("","",LevelThreshold(logrus.ErrorLevel)))

	logrus.Info("This is the info test message.")
	logrus.WithFields(SendTextMsg("This is the warn test message.",[]string{},false)).Warn()
	logrus.WithFields(SendMarkdownMsg("This is the warn test message send to dingtalk.","### Error Message",[]string{},false)).Error()
}