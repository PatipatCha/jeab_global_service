package model

type Project struct {
	ProjectId               int     `json:"id"`
	ProjectName             string  `json:"name"`
	ProjectLatitude         string  `json:"latitude"`
	ProjectLongitude        string  `json:"longitude"`
	ProjectDistanceLocation float32 `json:"distance_location"`
	ProjectHistoryAmount    int8    `json:"history_amount"`
	ProjectStatus           string  `json:"status"`
}
