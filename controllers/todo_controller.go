package controllers

import (
    "context"
    "onboard-clearisk/configs"
    "onboard-clearisk/models"
    "onboard-clearisk/responses"
    "net/http"
    "time"
  
    //"github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection = configs.GetCollection(configs.DB, "todos")
//var validate = validator.New()

func CreateTodo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var todo models.Todo
	defer cancel()

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&todo); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

	newTodo := models.Todo{
        Id:       primitive.NewObjectID(),
        Title:     todo.Title,
        Desc:      todo.Desc,
        Duration:  todo.Duration,
    }
	result, err := todoCollection.InsertOne(ctx, newTodo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.TodoResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllTodos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var todos []models.Todo
    defer cancel()
	results, err := todoCollection.Find(ctx, bson.M{})

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

	 //reading from the db in an optimal way
	defer results.Close(ctx)
	 for results.Next(ctx) {
		 var singleTodo models.Todo
		 if err = results.Decode(&singleTodo); err != nil {
			 return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		 }
 
		 todos = append(todos, singleTodo)
	 }
 
	 return c.Status(http.StatusOK).JSON(
		 responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": todos}},
	 )

}

func GetTodo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    todoId := c.Params("todoId")
    var todo models.Todo
    defer cancel()
  
    objId, _ := primitive.ObjectIDFromHex(todoId)
  
    err := todoCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&todo)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    return c.Status(http.StatusOK).JSON(responses.TodoResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": todo}})
}

func DeleteATodo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	todoId := c.Params("todoId")
	defer cancel()

	objId, _:= primitive.ObjectIDFromHex(todoId)
	result, err := todoCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
        return c.Status(http.StatusNotFound).JSON(
            responses.TodoResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Todo with specified ID not found!"}},
        )
    }

	return c.Status(http.StatusOK).JSON(
        responses.TodoResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Todo successfully deleted!"}},
    )
}

//func edit todos
//func edit ussers
func EditATodo(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    todoId := c.Params("todoId")
    var todo models.Todo
    defer cancel()
  
    objId, _ := primitive.ObjectIDFromHex(todoId)
  
    //validate the request body
    if err := c.BodyParser(&todo); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    //use the validator library to validate required fields
    if validationErr := validate.Struct(&todo); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.TodoResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }
  
    update := bson.M{"title": todo.Title, "desc": todo.Desc, "duration": todo.Duration}
  
    result, err := todoCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    //get updated user details
    var updatedTodo models.Todo
    if result.MatchedCount == 1 {
        err := todoCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedTodo)
        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.TodoResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
    }
  
    return c.Status(http.StatusOK).JSON(responses.TodoResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedTodo}})
}