package controllers

import (
	"github.com/PatipatCha/jeab_global_service/services"
	"github.com/gofiber/fiber/v2"
)

func GetConfigGlobalHandler(c *fiber.Ctx) (err error) {

	project_list, err := services.GetProjectLists()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	config_global_json := fiber.Map{
		"project_list": project_list,
	}

	return c.Status(200).JSON(config_global_json)
}

// var baseUrl = configuration.ApiProjectConfig()
// agent := fiber.Get(baseUrl)

// var project_list []model.ProjectModel

// statusCode, body, errs := agent.Bytes()
// if len(errs) > 0 {
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		"message": errs,
// 	})
// }

// err = json.Unmarshal(body, &project_list)
// if err != nil {
// 	log.Fatal(err)
// }

// jsonBytes, err := json.Marshal(project_list)
// if err != nil {
// 	log.Fatalf("Error marshaling struct: %v", err)
// }
// projectString := string(jsonBytes)

// project_list_gjson := gjson.Get(projectString, "#(status=\"active\")#")

// aa := project_list_gjson.String()

// byt := []byte(aa)

// err = json.Unmarshal(byt, &project_list)
// if err != nil {
// 	log.Fatal(err)
// }
