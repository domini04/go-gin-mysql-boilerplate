package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Validations"
)

func UserCreate(request *Validations.UserCreate) (err error) {
	if err = Config.DB.Table("users").Create(request).Error; err != nil {
		return err
	}
	return nil
}

func UserUpdate(request *Validations.UserUpdate, userId string) (err error) {
	if err = Config.DB.Table("users").Where("id = ?", userId).Update(request).Error; err != nil {
		return err
	}
	return nil
}