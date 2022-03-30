package hander

import (
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"session_manager/model"
	"strconv"
	"time"
)

const sessionStoreKey = "session_store"

var CreateSession gin.HandlerFunc = func(c *gin.Context) {
	token := uuid.NewString()
	sessionTimeout, _ := strconv.ParseInt(os.Getenv("SESSION_TIMEOUT"), 0, 64)
	expiration := time.Now().Add(time.Duration(sessionTimeout) * time.Second).UTC()
	encoder := scs.GobCodec{}
	sessionData := map[string]interface{}{}
	sessionDataEncoded, _ := encoder.Encode(expiration, sessionData)

	sessionStore := c.MustGet(sessionStoreKey).(*memstore.MemStore)
	err := sessionStore.Commit(token, sessionDataEncoded, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorOutput{
			Error: "could not create session",
		})
		return
	}

	c.JSON(http.StatusCreated, model.SessionOutput{
		Id:        token,
		ExpiresAt: expiration,
	})
	return
}

var GetSession gin.HandlerFunc = func(c *gin.Context) {
	sessionID := c.Param("id")
	sessionStore := c.MustGet(sessionStoreKey).(*memstore.MemStore)
	sessionData, found, err := sessionStore.Find(sessionID)

	if err != nil || found == false {
		c.JSON(http.StatusNotFound, model.ErrorOutput{
			Error: "session not found",
		})
		return
	}

	encoder := scs.GobCodec{}
	expiration, _, err := encoder.Decode(sessionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorOutput{
			Error: "failed to parse session data",
		})
		return
	}

	c.JSON(http.StatusOK, model.SessionOutput{
		Id:        sessionID,
		ExpiresAt: expiration,
	})
}

var DeleteSession gin.HandlerFunc = func(c *gin.Context) {
	sessionID := c.Param("id")
	sessionStore := c.MustGet(sessionStoreKey).(*memstore.MemStore)

	_, found, err := sessionStore.Find(sessionID)
	if err != nil || found == false {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "session not found",
		})
		return
	}

	if found == false {
		c.JSON(http.StatusOK, "")
		return
	}

	err = sessionStore.Delete(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete session",
		})
		return
	}

	c.JSON(http.StatusOK, "")
}
