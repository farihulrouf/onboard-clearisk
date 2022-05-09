package main

import (
    "BE_CLEARISK.IO/configs" 
    "github.com/gofiber/fiber/v2" 
)

func main() {
    app := fiber.New()

    //run database
    configs.ConnectDB()

    app.Listen("127.0.0.1:2222")
}

/*
app.Listen(8080)
app.Listen("8080")
app.Listen(":8080")
app.Listen("127.0.0.1:8080")
*/