package public

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// 项目初始化
func init() {
	initLog()
	initConf()
}

// initLog 初始化日志
func initLog() {
	// 设置日志输出格式
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 Monday"
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// 创建日志保存目录
	rootdir,err := os.Getwd()
	if err != nil {
		log.Debug().Str("error",err.Error()).Msg("获取当前目录地址失败")
	}
	logPath := fmt.Sprintf("%s/doc/log",rootdir)
	// 判断当前文件夹是否存在
	if _,err = os.Stat(logPath);err != nil {
		err = os.MkdirAll(logPath,os.ModePerm)
		log.Debug().Str("error",err.Error()).Msg("创建日志存储目录失败")
	}
	// 创建日志文件
	logFile := fmt.Sprintf("%s/%s.log",logPath,time.Now().Format("20060102"))
	f,err := os.OpenFile(logFile,os.O_APPEND|os.O_APPEND|os.O_WRONLY,os.ModePerm)
	if err != nil {
		log.Debug().Str("error",err.Error()).Msg("创建日志存文件失败")
	}

	// 更新日志输出格式
	logout := zerolog.ConsoleWriter{Out: os.Stdout}
	mulout := zerolog.MultiLevelWriter(logout, f)
	log.Logger = zerolog.New(mulout).With().Timestamp().Caller().Logger()
}

// initConf 初始化配置文件
func initConf() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf/")

	viper.ReadInConfig()
	viper.WatchConfig()
}