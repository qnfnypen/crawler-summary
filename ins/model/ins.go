
package model

// InsInfo 页面响应的json
// 使用http://json2struct.mervine.net/生成json
type InsInfo struct {
	Data struct {
		User struct {
			EdgeOwnerToTimelineMedia struct {
				Edges []struct {
					Node struct {
						DisplayURL string `json:"display_url"`
						IsVideo    bool   `json:"is_video"`
						VideoURL   string `json:"video_url"`
					} `json:"node"`
				} `json:"edges"`
				PageInfo struct {
					EndCursor   string `json:"end_cursor"`
					HasNextPage bool   `json:"has_next_page"`
				} `json:"page_info"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"data"`
}