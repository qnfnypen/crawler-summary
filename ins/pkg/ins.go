package pkg

// GetInsURL 获取Ins上的图片和视频链接
func GetInsURL() {
	// id := getUserID()
	// fmt.Println(id)

	// var after string

	// // for {
	// 	url := joinURL(id, after)
	// 	var ins model.InsInfo
	// 	content := getURLContent(url)
	// 	json.Unmarshal([]byte(content), &ins)
	// 	for _, v := range ins.Data.User.EdgeOwnerToTimelineMedia.Edges {
	// 		if v.Node.IsVideo {
	// 			util.WriteVideoToFile(v.Node.VideoURL)
	// 		} else {
	// 			util.WriteImgToFile(v.Node.DisplayURL)
	// 		}
	// 	}
	// 	// 判断是否是最后一页
	// 	if !ins.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage {
	// 		return
	// 	}

	// 	after = getAfter(content)
	// // }
}
