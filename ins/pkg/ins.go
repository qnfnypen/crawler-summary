package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/qnfnypen/crawler-summary/ins/model"
	"github.com/qnfnypen/crawler-summary/ins/util"
)

// GetInsURL 获取Ins上的图片和视频链接
func GetInsURL() {
	id := getUserID()

	var after string

	for {
		// 拼接内存url
		url := joinURL(id, after)
		fmt.Println(url)
		var ins model.InsInfo
		// 获取页面内容
		content := getURLContent(url)
		json.Unmarshal([]byte(content), &ins)
		for _, v := range ins.Data.User.EdgeOwnerToTimelineMedia.Edges {
			if v.Node.IsVideo {
				util.WriteVideoToFile(v.Node.VideoURL)
			} else {
				util.WriteImgToFile(v.Node.DisplayURL)
			}
		}
		// 判断是否是最后一页
		if !ins.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage {
			return
		}

		after = getAfter(content)
	}
}
