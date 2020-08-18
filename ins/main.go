package main

import (
	"fmt"

	"github.com/spf13/viper"

	// 初始化项目
	_ "github.com/qnfnypen/crawler-summary/ins/public"
)

func main() {
	fmt.Println(viper.GetString("HTTP.URL"))
	fmt.Println(viper.GetString("HTTP.QUERY_HASH"))
}
