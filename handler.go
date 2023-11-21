package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PatipatCha/jeab_global_service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupApiRoutes(app)

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	err := app.Listen(listenAddr)
	if err != nil {
		log.Fatalf("Error while starting Fiber: %v", err)
	}

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
