package handler

import (
	"converter/model"
	"converter/romantoarabic"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ConvertRomanNumeral gin.HandlerFunc = func(c *gin.Context) {
	var json model.ConvertRomanNumeral

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorOutput{
			Error: err.Error(),
		})
		return
	}

	output, err := romantoarabic.ToInteger(json.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorOutput{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"output": output})
}
