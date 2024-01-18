package adapters

import (
	entites "github.com/PatipatCha/jeab_global_service/entities"
	"github.com/PatipatCha/jeab_global_service/model"
)

func MapEntityToModel(entities []entites.ProjectEntity) []model.ProjectModel {
	var data []model.ProjectModel

	for _, proj := range entities {
		project := model.ProjectModel{
			SEOCID:                  proj.SEOCID,
			ProjectId:               proj.OperationCenterUserID,
			ProjectNameTH:           proj.NameTH,
			ProjectNameEN:           proj.NameEN,
			ProjectLatitude:         proj.Latitude,
			ProjectLongitude:        proj.Longitude,
			ProjectDistanceLocation: 0.1,
		}

		data = append(data, project)
	}

	return data
}
