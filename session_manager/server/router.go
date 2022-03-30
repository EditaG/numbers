package server

import (
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/gin-gonic/gin"
	"net/http"
	"session_manager/hander"
	"time"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.Use(getSessionStoreMiddleware(getSessionStore()))

	router.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, "OK")
	})

	sessionRoutes := router.Group("/session")
	{
		sessionRoutes.POST("/", hander.CreateSession)
		sessionRoutes.GET("/:id", hander.GetSession)
		sessionRoutes.DELETE("/:id", hander.DeleteSession)
	}

	router.NoRoute(func(context *gin.Context) { context.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func getSessionStore() *memstore.MemStore {
	return memstore.NewWithCleanupInterval(30 * time.Second)
}

func getSessionStoreMiddleware(store *memstore.MemStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("session_store", store)
		c.Next()
	}
}
