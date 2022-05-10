
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

    app.Get("/todos", controllers.GetAllTodos)
    app.Post("/todo", controllers.CreateTodo)
    app.Get("/todo/:todoId", controllers.GetTodo)
    app.Delete("/todo/:todoId", controllers.DeleteATodo)
    app.Put("/todo/:todoId", controllers.EditATodo)
}