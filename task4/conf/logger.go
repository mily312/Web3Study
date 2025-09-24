package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zap.InfoLevel
	}
	zapCore := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)

	return zap.New(zapCore).Sugar()
}

// 定义输出日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// 定义日志文件路径，保留天数，最大文件大小。。。
func getWriteSyncer() zapcore.WriteSyncer {
	//定义生成文件路径
	stWorkDir, _ := os.Getwd()
	stSeparator := string(filepath.Separator)
	//stRootDir := stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]
	stLogFilePath := stWorkDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"

	fmt.Println("stLogFilePath:" + stLogFilePath)

	//使用lumberjack写入日志
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"), // megabytes
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"), //days
		Compress:   false,                      // disabled by default
	}

	return zapcore.AddSync(lumberjackSyncer)
}
