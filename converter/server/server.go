package server

import "github.com/joho/godotenv"

func Start() {
	initEnv()
	router := setRouter()
	router.Run(":8080")
}

func initEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}
