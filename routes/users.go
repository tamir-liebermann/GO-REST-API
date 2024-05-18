package routes

import (
	"net/http"

	"example.com/REST-API/models"
	"example.com/REST-API/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

    err := context.ShouldBindJSON(&user) 

  if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"massage": "Could not parse request data." } )
	return 
  }

 err = user.Save()

 if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"massage": "Could not save user." } )
	return 
  }

  context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func login(context *gin.Context) {
	var user models.User

    err := context.ShouldBindJSON(&user) 

    if err != nil {
	 context.JSON(http.StatusBadRequest, gin.H{"massage": "Could not parse request data." } )
	 return 
    }

	err = user.ValidateCredatials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Could not authenticate user."})
		return 
	}

 token , err := utils.GenerateToken(user.Email, user.ID)
 
    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not authenticate user."})
		return 
	}

    context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}