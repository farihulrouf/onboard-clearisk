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
var postCollection *mongo.Collection = configs.GetCollection(configs.DB, "post")
func CreatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var post models.Post
	defer cancel()

	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DataResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&post); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.DataResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

	newPost := models.Post{
        Id:       primitive.NewObjectID(),
        Title:     post.Title,
        Content:      post.Content,
    }
	result, err := postCollection.InsertOne(ctx, newPost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.DataResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllPost(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var posts []models.Post
    defer cancel()

    results, err := userCollection.Find(ctx, bson.M{})

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singlePost models.Post
        if err = results.Decode(&singlePost); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }

        posts = append(posts, singlePost)
    }

    return c.Status(http.StatusOK).JSON(
        responses.DataResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": posts}},
    )
}


