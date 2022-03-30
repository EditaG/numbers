package handler

import (
	"converter/model"
	"converter/romantoarabic"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CheckAuth gin.HandlerFunc = func(c *gin.Context) {
	var json model.ConvertRomanNumeral

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := romantoarabic.ToInteger(json.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"output": output})
}
