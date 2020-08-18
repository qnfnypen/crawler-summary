package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	// 使用日志和配置
	_ "github.com/qnfnypen/crawler-summary/ins/public"
)

// GetURLContent 获取页面内容
func GetURLContent(urlStr string) string {
	// 解析代理地址
	proxyAddr := viper.GetString("URL_PROXY")
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal().Str("error", err.Error()).Msg("解析代理地址失败")
	}
	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		ResponseHeaderTimeout: 5 * time.Second,
		MaxIdleConnsPerHost:   10,
	}

	// 创建客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
	fmt.Println(urlStr)
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Fatal().Str("error", err.Error()).Msg("创建请求连接失败")
	}
	req.Header.Add("user-agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36`)
	req.Header.Set("cookie", viper.GetString("Cookie"))
	resp, err := client.Do(req)
	if err != nil  {
		log.Fatal().Str("error", err.Error()).Msg("获取页面内容失败")
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal().Str("error", "状态码不为200").Msg("获取页面内容失败")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Str("error", err.Error()).Msg("读取页面内容失败")
	}

	return string(respBody)
}

// getUserID 获取用户参数，进行内容拼接
// func getUserID() string {
// 	resp := getURLContent(viper.GetString("URL"))
// 	regex := `{"id":"(.*?)","username":`
// 	rp := regexp.MustCompile(regex)
// 	ss := rp.FindStringSubmatch(resp)

// 	return ss[1]
// }

// getAfter 获取下一页的文档游标
func getAfter(content string) string {
	// "end_cursor":"(.*)"},"edges":
	regex := `"end_cursor":"(.*?)"},"edges":`
	rp := regexp.MustCompile(regex)
	ss := rp.FindStringSubmatch(content)

	return ss[1]
}

// joinURL 对参数进行拼接
func joinURL(id, after string) string {
	if after == "" {
		return fmt.Sprintf(`https://www.instagram.com/graphql/query/?query_hash=%s&variables={"id":"%s","first":12}`,
			viper.GetString("QUERY_HASH"), id)
	}

	return fmt.Sprintf(`https://www.instagram.com/graphql/query/?query_hash=%s&variables={"id":"%s","first":12,"after":"%s"}`,
		viper.GetString("QUERY_HASH"), id, after)
}
