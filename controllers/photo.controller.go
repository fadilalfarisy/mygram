package controllers

import (
	"challenge-4/databases"
	"challenge-4/helpers"
	"challenge-4/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto 		godoc
// @Summary 			create photo
// @Description 	create photo user
// @Tags 					photo
// @Accept 				application/json
// @Produce 			application/json
// @Param 				models.Photo body models.Photo true "photo user"
// @Success 			200 {object} models.Photo
// @Router 				/photo/ [post]
func CreatePhoto(c *gin.Context) {
	db := databases.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// GetAllPhoto 	godoc
// @Summary 			get all photo
// @Description 	get all photo users
// @Tags 					photo
// @Accept 				application/json
// @Produce 			application/json
// @Success 			200 {object} models.Photo
// @Router 				/photo/ [get]
func GetAllPhoto(c *gin.Context) {
	db := databases.GetDB()

	Photo := []models.Photo{}

	err := db.Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": Photo,
	})
}

// GetPhotoById 	godoc
// @Summary 			get photo by id
// @Description 	get photo user by id
// @Tags 					photo
// @Accept 				application/json
// @Produce 			application/json
// @Param					photoId path int true "id of the photo"
// @Success 			200 {object} models.Photo
// @Router 				/photo/{photoId} [get]
func GetPhotoById(c *gin.Context) {
	db := databases.GetDB()

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.First(&Photo, "id = ?", photoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// UpdatePhoto 	godoc
// @Summary 			update photo
// @Description 	update photo user
// @Tags 					photo
// @Accept 				application/json
// @Produce 			application/json
// @Param					photoId path int true "id of the photo"
// @Success 			200 {object} models.Photo
// @Router 				/photo/{photoId} [put]
func UpdatePhoto(c *gin.Context) {
	db := databases.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		Photo_Url: Photo.Photo_Url,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto 	godoc
// @Summary 			delete photo
// @Description 	delete photo user
// @Tags 					photo
// @Accept 				application/json
// @Produce 			application/json
// @Param					photoId path int true "id of the photo"
// @Success 			200 {object} models.Photo
// @Router 				/photo/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Photo)
}
