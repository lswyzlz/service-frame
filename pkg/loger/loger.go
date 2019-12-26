package loger

import (
	"os"
	"path"
	"service-frame/pkg/common"
	"service-frame/pkg/setting"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	//Apploger 日志对象
	Apploger *logrus.Logger
)

func init() {
	levelMap := make(map[string]logrus.Level, 7)
	levelMap["PanicLevel"] = logrus.PanicLevel
	levelMap["FatalLevel"] = logrus.FatalLevel
	levelMap["ErrorLevel"] = logrus.ErrorLevel
	levelMap["WarnLevel"] = logrus.WarnLevel
	levelMap["InfoLevel"] = logrus.InfoLevel
	levelMap["DebugLevel"] = logrus.DebugLevel
	levelMap["TraceLevel"] = logrus.TraceLevel

	Apploger = logrus.New()

	//日志输出格式
	Apploger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})

	//设置日志级别
	if sL := setting.AppConfig.GetString("log.level"); sL == "" {
		if setting.AppConfig.GetString("runmode") == "release" {
			Apploger.SetLevel(logrus.InfoLevel)
		}
	} else {
		tL, ok := levelMap[sL]
		if !ok {
			Apploger.WithField("读取的日志级别", sL).Fatalln("日志级别配置错误")
		}
		Apploger.SetLevel(tL)
	}

	//设置输出位置
	if sName := setting.AppConfig.GetString("log.filename"); sName == "" {
		Apploger.SetOutput(os.Stdout)
	} else {
		//设置默认参数
		setting.AppConfig.SetDefault("log.filepath", "./log")
		setting.AppConfig.SetDefault("log.maxage", 24*7)
		setting.AppConfig.SetDefault("log.rotation", 24)

		sPath := setting.AppConfig.GetString("log.filepath")
		iMaxage := time.Duration(setting.AppConfig.GetInt32("log.maxage")) * time.Hour
		iRotation := time.Duration(setting.AppConfig.GetInt32("log.rotation")) * time.Hour

		Apploger.AddHook(newRotateHook(sPath, sName, iMaxage, iRotation))
	}

	//设置没有锁
	//Apploger.SetNoLock()
}

//newRotateHook 日志文件分割Hook
func newRotateHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	appPath, err := common.GetAppPath()
	if err != nil {
		Apploger.WithField("错误信息", err.Error()).Fatalln("获取运行目录失败")
	}

	baseLogPath := path.Join(appPath, logPath)
	err = common.MakeDir(baseLogPath)
	if err != nil {
		Apploger.WithField("错误信息", err.Error()).Fatalln("创建日志目录失败")
	}

	baseLogPath = path.Join(baseLogPath, logFileName)

	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		Apploger.WithField("错误信息", err.Error()).Fatalln("配置日志分割错误")
	}

	Apploger.SetOutput(writer)
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, Apploger.Formatter)
}
