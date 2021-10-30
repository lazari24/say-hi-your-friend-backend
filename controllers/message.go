package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"say-hi-backend/config"
	"say-hi-backend/models"
	"time"
)

func SaveMessage(c *fiber.Ctx) error {
	dbCollection := config.DatabaseInstance.DB.Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	message := new(models.Message)

	if err := c.BodyParser(message); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error": err,
		})
	}

	result, err := dbCollection.InsertOne(ctx, message)

	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert Message",
			"error": err,
		})
	}

	return c.Status(200).JSON(result)
}

func GetOneMessageById(c *fiber.Ctx) error {
	dbCollection := config.DatabaseInstance.DB.Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	message := new(models.Message)

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := dbCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrase Not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&message)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrase Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(message)
}
