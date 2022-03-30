package handler

import (
	"converter/auth"
	"converter/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var Login gin.HandlerFunc = func(c *gin.Context) {
	var request model.Login

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorOutput{
			Error: err.Error(),
		})

		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	m := model.User{}
	user, err := m.FindUserByUsername(db, request.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorOutput{
			Error: "invalid credentials",
		})

		return
	}

	err = auth.VerifyPassword(user.Password, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorOutput{
			Error: "invalid credentials",
		})

		return
	}

	session, err := auth.CreateSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorOutput{
			Error: "failed to initialize session",
		})

		return
	}

	c.JSON(http.StatusOK, session)
}
