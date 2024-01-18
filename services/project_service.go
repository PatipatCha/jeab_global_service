package services

import (
	"github.com/PatipatCha/jeab_global_service/adapters"
	"github.com/PatipatCha/jeab_global_service/model"
)

func GetProjectLists() ([]model.ProjectModel, error) {
	entity := GetProjectApi()
	res := adapters.MapEntityToModel(entity)

	return res, nil
}
