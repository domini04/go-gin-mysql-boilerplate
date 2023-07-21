package Controllers

import (
	"go-gin-mysql-boilerplate/Models"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Services"
	"go-gin-mysql-boilerplate/Validations"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserFetchAll(c *gin.Context) {
	var user []Schema.User
	err := Models.UserFetchAll(&user)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, nil, user)
	}
}

func UserFetchSingle(c *gin.Context) {
	userId := c.Param("id")

	var user Schema.User
	err := Models.UserFetchSingle(&user, userId)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, nil, user)
	}
}

func UserCreate(c *gin.Context) {
	var request Validations.UserCreate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	request.Password = Services.MD5Hash(request.Password)
	saveErr := Models.UserCreate(&request)
	if saveErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", saveErr)
	} else {
		Services.Success(c, "Created", request)
	}
}

func UserUpdate(c *gin.Context) {
	userId := c.Param("id")

	var request Validations.UserUpdate
	if requestErr := c.ShouldBind(&request); requestErr != nil {
		errRes := strings.Split(requestErr.Error(), ": ")
		Services.ValidationError(c, "These fields are required!", errRes)
		return
	}

	updateErr := Models.UserUpdate(&request, userId)
	if updateErr != nil {
		Services.NotAcceptable(c, "Something went wrong!", updateErr)
	} else {
		Services.Success(c, "Updated", request)
	}
}

func UserDelete(c *gin.Context) {
	userId := c.Param("id")

	var user Schema.User
	err := Models.UserDelete(&user, userId)
	if err != nil {
		Services.NotAcceptable(c, "Something went wrong!", err)
	} else {
		Services.Success(c, "Deleted", nil)
	}
}

// Create DogFetchAll api that returns image from https://storage.googleapis.com/cdn-bucket-mz-finalproject/doge.jpg
func DogFetchAll(c *gin.Context) {
	Services.Success(c, nil, "https://storage.googleapis.com/cdn-bucket-mz-finalproject/doge.jpg")
}
