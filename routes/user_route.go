
package routes

import (
    "onboard-clearisk/controllers" //add this

    "github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
    app.Post("/user", controllers.CreateUser)
    app.Get("/users", controllers.GetAllUsers)
    app.Get("/user/:userId", controllers.GetAUser)
    app.Put("/user/:userId", controllers.EditAUser)
    app.Delete("/user/:userId", controllers.DeleteAUser)
}