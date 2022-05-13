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
    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)
var jobdeskCollection *mongo.Collection = configs.GetCollection(configs.DB, "jobdesk")
func CreateJobdesk(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var jobdesk models.Jobdesk
	defer cancel()

	if err := c.BodyParser(&jobdesk); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DataResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&jobdesk); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.DataResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

	newJobdesk := models.Jobdesk{
        Id:       primitive.NewObjectID(),
        Name:     jobdesk.Name,
        Desc:      jobdesk.Desc,
        Duration:  jobdesk.Duration,
    }
	result, err := jobdeskCollection.InsertOne(ctx, newJobdesk)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.DataResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}


