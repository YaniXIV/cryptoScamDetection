package handlers

import (
	"cryptoScamDetection/models"
	"cryptoScamDetection/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleAnalyze(c *gin.Context) {
	var bodyData models.ResponseData
	if bindErr := c.ShouldBindJSON(&bodyData); bindErr != nil {
		c.AbortWithError(http.StatusBadRequest, bindErr)
	}
	//After this we are going to add our llm api call and logic.
	generatedResponse := services.GenResponse(bodyData.Data)

	c.JSON(http.StatusAccepted, gin.H{
		"message":           "Data received!",
		"GeneratedResponse": generatedResponse,
	})
}
