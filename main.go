package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server.hrszpuk.com/api"
)

func main() {
	log.Printf("Setting up server")

	log.Printf("Doing preemptive checks")
	DirectoryStatusCheck(api.ContainersDir)
	DirectoryStatusCheck(api.ScriptsDir)

	app := fiber.New()

	log.Printf("Setting up routes")

	// Server website routes
	app.Get("/", home)

	// Server API routes
	apiGroup := app.Group("api")
	v1 := apiGroup.Group("v1")

	v1.Get("ping", api.Ping)
	v1.Get("uptime", api.Uptime)
	v1.Get("system", api.System)

	containers := v1.Group("containers")

	containers.Get("list", api.ContainersList)
	containers.Get("get/:id", api.ContainersGet)

	scripts := v1.Group("scripts")

	scripts.Get("list", api.ScriptsList)
	scripts.Get("run/:id", api.ScriptsRun)

	log.Printf("Starting server. Listening on port 3000.")
	log.Fatal(app.Listen(":3000"))
}

func home(ctx *fiber.Ctx) error {
	err := ctx.SendFile("pages/home.html")
	if err == nil {
		log.Printf("Served home.html to %s", ctx.IP())
	} else {
		log.Printf("An error occurred while trying to serve home.html to %s. Error: %s", ctx.IP(), err.Error())
	}
	return nil
}
