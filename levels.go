package dingtalkrus

import "github.com/sirupsen/logrus"

// Supported log levels
var AllLevels=[]logrus.Level{
	logrus.WarnLevel,
	logrus.ErrorLevel,
	logrus.FatalLevel,
	logrus.PanicLevel,
	logrus.DebugLevel,
}

// Returns every logging level above and including the given parameter.
func LevelThreshold(l logrus.Level)[]logrus.Level  {
	for i:=range AllLevels {
		if AllLevels[i]==l{
			return AllLevels[i:]
		}
	}
	return []logrus.Level{}
}