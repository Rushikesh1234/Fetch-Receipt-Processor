package handlers

import (
	"net/http"
	"sync"

	"receipt-processor/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receiptsStore = make(map[string]map[string]interface{})
var mutex = &sync.Mutex{}

func ProcessReceipt(c *gin.Context) {
	var receipt map[string]interface{}
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	receiptID := uuid.New().String()
	points := utils.CalculatePoints(receipt)

	mutex.Lock()
	receiptsStore[receiptID] = map[string]interface{}{
		"receipt": receipt,
		"points":  points,
	}
	mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

func GetPoints(c *gin.Context) {
	receiptID := c.Param("id")

	mutex.Lock()
	receiptData, exists := receiptsStore[receiptID]
	mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	points := receiptData["points"].(int)
	c.JSON(http.StatusOK, gin.H{"points": points})
}
