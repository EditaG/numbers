package model

type Login struct {
	Username string `binding:"required,max=255"`
	Password string `binding:"required,max=255"`
}
