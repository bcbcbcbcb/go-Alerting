package utils

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Logger = &lumberjack.Logger{
	Filename:   "./logs/alert.txt", // 日志文件得位置
	MaxSize:    100,                // 切割之前日志文件得大小（单位：MB）
	MaxBackups: 0,                  // 保留旧文件得最大个数; 0 表示不限制文件个数
	MaxAge:     28,                 // 保留旧文件得最大天数
	Compress:   true,               // 是否压缩旧文件
	LocalTime:  true,               // 是否使用本地时间；默认 UTC 时间
}

var Log = logrus.New()

func init() {
	// logger.Formatter = new(logrus.JSONFormatter) 				// 输出格式为json
	Log.Formatter = new(logrus.TextFormatter)                      // 输出格式为txt文本
	Log.Formatter.(*logrus.TextFormatter).DisableColors = true     // remove colors
	Log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // remove timestamp from test output
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
	})

	// 日志输出最低级别
	Log.Level = logrus.TraceLevel

	// 输出日志位置
	// log.Out = os.Stdout // 默认输出
	// logfile, err := os.OpenFile("alert.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = logfile
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }
	Log.Out = Logger // 输出至创建的 lumberjack 的 logger 中

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			Log.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()

	// 使用输出日志到文件：级别日志 info、warn、debug、panic、trace 输出示例：
	// log.WithFields(logrus.Fields{"animal": "walrus", "number": 0}).Trace("Went to the beach")
	// log.WithFields(logrus.Fields{"animal": "walrus", "number": 8}).Debug("Started observing beach")
	// log.WithFields(logrus.Fields{"animal": "walrus", "size": 10}).Info("A group of walrus emerges from the ocean")
	// log.WithFields(logrus.Fields{"omg": true, "number": 122}).Warn("The group's number increased tremendously!")
	// log.WithFields(logrus.Fields{"temperature": -4}).Debug("Temperature changes")
	// log.WithFields(logrus.Fields{"animal": "orca", "size": 9009}).Panic("It's over 9000!")
}
