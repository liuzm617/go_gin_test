package models

import "github.com/jinzhu/gorm"

type Auth struct {
	Model
	Username string `gorm:"type:varchar(50);unique_index" json:"username"`
	Password string `json:"password"`

}

// check user exists

func CheckAuth(username, password string)(bool, error){
	var user Auth
	err := db.Select("id").Where(Auth{Username:username, Password:password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}
	if user.ID > 0{
		return true, nil
	}
	return false, nil
}

