package handler

import (
	"converter/auth"
	"converter/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Logout gin.HandlerFunc = func(c *gin.Context) {
	sessionDeleted, _ := auth.DeleteSession(auth.GetAuthToken(c))
	if sessionDeleted == false {
		c.JSON(http.StatusInternalServerError, model.ErrorOutput{
			Error: "failed to delete session",
		})

		return
	}

	c.JSON(http.StatusOK, "")
}
