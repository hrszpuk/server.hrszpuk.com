package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	v1 := app.Group("v1")

	v1.Get("docs", docs)
	v1.Get("ping", ping)
	v1.Get("uptime", uptime)
	v1.Get("os", os)

	containers := v1.Group("containers")

	containers.Get("list", containersList)
	containers.Get("get/:id", containersGet)

	scripts := v1.Group("scripts")
	scripts.Get("list", scriptsList)
	scripts.Get("run/:id", scriptsRun)

	log.Fatal(app.Listen(":3000"))
}
