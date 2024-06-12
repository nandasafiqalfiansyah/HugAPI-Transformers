package main

import (
	"huging-face/model"
	"huging-face/repository"
	"huging-face/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize repository
    apiKey := os.Getenv("HUGGING_FACE_API_KEY")
    huggingFaceRepo := repository.NewHuggingFaceRepository(apiKey)

    // Initialize service
    inferenceService := service.NewInferenceService(huggingFaceRepo)

    // Initialize Gin router
    router := gin.Default()

    // Endpoint handler for inference
    router.POST("/inference", func(c *gin.Context) {
		var req model.Payload
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        result, err := inferenceService.Inference(model.Payload{
			Inputs: req.Inputs,})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
            return
        }

        c.JSON(http.StatusOK, result)
    })

    // Run the server
    log.Fatal(router.Run(":8080"))
}