package controllers

import (
	"challenge-4/databases"
	"challenge-4/helpers"
	"challenge-4/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister	godoc
// @Summary 		User register
// @Description Create account user
// @Tags 				user
// @Accept 			application/json
// @Produce 		application/json
// @Param 			models.User body models.User true "register user"
// @Success 		200 {object} models.User
// @Router 			/user/register [post]
func UserRegister(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

// UserLogin 			godoc
// @Summary 			User login
// @Description 	Login account user
// @Tags 					user
// @Accept 				application/json
// @Produce 			application/json
// @Param 				models.User body models.User true "login user"
// @Success 			200 {object} models.User
// @Router 				/user/login [post]
func UserLogin(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unanuthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unanuthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
