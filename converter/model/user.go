package model

import (
	"converter/auth"
	"gorm.io/gorm"
)

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := auth.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	_ = u.BeforeSave()
	err := db.Create(&u).Error

	return u, err
}

func (u *User) FindUserByUsername(db *gorm.DB, username string) (*User, error) {
	var err error
	err = db.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, err
}

func GenerateDemoUser(db *gorm.DB) error {
	firstUser := User{
		Username: "demo_user",
		Password: "Pa33m0rD*&!",
	}

	_, err := firstUser.FindUserByUsername(db, firstUser.Username)
	if err != nil {
		_, err = firstUser.SaveUser(db)
		if err != nil {
			return err
		}
	}

	return nil
}
