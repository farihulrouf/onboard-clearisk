package main

import (
    "onboard-clearisk/configs"
    "onboard-clearisk/routes" //add this
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    //run database
    configs.ConnectDB()

    //routes
    routes.UserRoute(app) //add this

    app.Listen(":2222")
}