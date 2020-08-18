package util

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// getUserName 获取用户名，用户创建文件夹
func getUserName() string {
	url := viper.GetString("URL")
	regex := `://(.*)`
	rp := regexp.MustCompile(regex)
	ss := rp.FindStringSubmatch(url)
	s := strings.Split(ss[1], "/")

	return s[len(s)-2]
}

// WriteImgToFile 将图片链接写入文件
func WriteImgToFile(url string) {
	// 获取用户名创建文件夹
	name := getUserName()
	fileDir := fmt.Sprintf("doc/%s", name)
	if _, err := os.Stat(fileDir); err != nil {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatal().Str("error", err.Error()).Msg("创建存储文件夹失败")
		}
	}
	filePath := fmt.Sprintf("%s/img.txt", fileDir)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal().Str("error", err.Error()).Msg("创建存储文件失败")
	}

	f.WriteString(url + "\n")
}

// WriteVideoToFile 将视频链接写入文件
func WriteVideoToFile(url string) {
	// 获取用户名创建文件夹
	name := getUserName()
	fileDir := fmt.Sprintf("doc/%s", name)
	if _, err := os.Stat(fileDir); err != nil {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatal().Str("error", err.Error()).Msg("创建存储文件夹失败")
		}
	}
	filePath := fmt.Sprintf("%s/video.txt", fileDir)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal().Str("error", err.Error()).Msg("创建存储文件失败")
	}

	f.WriteString(url + "\n")
}
