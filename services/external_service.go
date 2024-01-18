package services

import (
	"encoding/json"

	"github.com/PatipatCha/jeab_global_service/configuration"
	entites "github.com/PatipatCha/jeab_global_service/entities"
	"github.com/gofiber/fiber/v2"
)

func GetProjectApi() []entites.ProjectEntity {
	var baseUrl = configuration.ApiProjectConfig()
	agent := fiber.Get(baseUrl)

	var entity []entites.ProjectEntity

	_, body, _ := agent.Bytes()

	json.Unmarshal(body, &entity)

	return entity
}

// jsonBytes, err := json.Marshal(project_list)
// if err != nil {
// 	return project, err
// }

// projectString := string(jsonBytes)

// project_list_gjson := gjson.Get(projectString, "#(status=\"active\")#")

// project_active := project_list_gjson.String()

// byt := []byte(project_active)

// err = json.Unmarshal(byt, &project_list)
// if err != nil {
// 	log.Fatal(err)
// }
