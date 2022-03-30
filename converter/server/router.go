package server

import (
	"converter/auth"
	"converter/handler"
	"converter/model"
	validators "converter/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var DB *gorm.DB

func setRouter() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	registerValidators()
	router.Use(getDatabaseMiddleware())

	router.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, "OK")
	})

	api := router.Group("/api")
	{
		api.POST("login", handler.Login)
		api.POST("logout", getAuthenticationMiddleware(), handler.Logout)
		api.POST("convert/roman", getAuthenticationMiddleware(), handler.ConvertRomanNumeral)
	}

	router.NoRoute(func(context *gin.Context) { context.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is_roman_numeral", validators.IsRomanNumeral)
	}
}

func getAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if auth.AuthTokenValid(c) == false {
			c.JSON(http.StatusUnauthorized, "")
			c.Abort()
		}
	}
}

func getDatabaseMiddleware() gin.HandlerFunc {
	database, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")))
	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	database.AutoMigrate(&model.User{})
	model.GenerateDemoUser(database)

	DB = database

	return func(c *gin.Context) {
		c.Set("DB", DB)
		c.Next()
	}
}
