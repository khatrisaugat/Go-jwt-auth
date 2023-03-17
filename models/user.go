package models

import (
	"github.com/khatrisaugat/JWTPractise/helpers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null;unique"`
}

func verifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	DB := helpers.ConnectDB()
	u := User{}
	err = DB.Model(User{}).Where("username=?", username).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = verifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {

		return "", err
	}
	token, err := helpers.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) SaveUser() (*User, error) {
	DB := helpers.ConnectDB()
	DB.AutoMigrate(&User{})
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func GetUserByID(user_id uint) (User, error) {
	DB := helpers.ConnectDB()
	u := User{}
	err := DB.Model(User{}).First(&u, user_id).Error
	if err != nil {
		return u, err
	}
	return u, nil
}
