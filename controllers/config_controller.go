package controllers

import (
	"encoding/json"
	"log"

	"github.com/PatipatCha/jeab_global_service/configuration"
	"github.com/PatipatCha/jeab_global_service/model"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetConfigGlobalHandler(c *fiber.Ctx) (err error) {
	var baseUrl = configuration.ApiProjectConfig()
	agent := fiber.Get(baseUrl)

	var project_list []model.Project

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errs,
		})
	}

	err = json.Unmarshal(body, &project_list)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.Marshal(project_list)
	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}
	projectString := string(jsonBytes)

	project_list_gjson := gjson.Get(projectString, "#(status=\"active\")#")

	aa := project_list_gjson.String()

	byt := []byte(aa)

	err = json.Unmarshal(byt, &project_list)
	if err != nil {
		log.Fatal(err)
	}

	config_global_json := fiber.Map{
		"project_list": project_list,
	}

	return c.Status(statusCode).JSON(config_global_json)
}
