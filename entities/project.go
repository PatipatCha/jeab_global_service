package entites

import "time"

type ProjectEntity struct {
	JeabID                string    `json:"jeab_id"`
	Username              string    `json:"username"`
	Role                  string    `json:"role"`
	Status                string    `json:"status"`
	CreatedBy             string    `json:"created_by"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	OperationCenterUserID int       `json:"operation_center_user_id"`
	SEOCID                int       `json:"seoc_id"`
	ProfileJeabID         string    `json:"profile_jeab_id"`
	NameTH                string    `json:"name_th"`
	NameEN                string    `json:"name_en"`
	Mobile                string    `json:"mobile"`
	ImageLogo             string    `json:"image_logo"`
	ProvinceID            string    `json:"province_id"`
	AmphurID              string    `json:"amphur_id"`
	TambonID              string    `json:"tambon_id"`
	ZipCode               string    `json:"zip_code"`
	Latitude              string    `json:"latitude"`
	Longitude             string    `json:"longitude"`
}
