package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
  
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB Siplah"})
    })
    app.Listen("127.0.0.1:2222")
    ///app.Listen(":8111")
    //app.Listen(8282)
    //router.Run(":8111")
}

/*
app.Listen(8080)
app.Listen("8080")
app.Listen(":8080")
app.Listen("127.0.0.1:8080")
*/