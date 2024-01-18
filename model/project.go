package model

type ProjectModel struct {
	ProjectId               int     `json:"id"`
	ProjectNameTH           string  `json:"name_th"`
	ProjectNameEN           string  `json:"name_en"`
	ProjectLatitude         string  `json:"latitude"`
	ProjectLongitude        string  `json:"longitude"`
	ProjectDistanceLocation float32 `json:"distance_location"`
	// ProjectHistoryAmount    int8    `json:"history_amount"`
	ProjectStatus string `json:"status"`
	SEOCID        int    `json:"seoc_id"`
}
