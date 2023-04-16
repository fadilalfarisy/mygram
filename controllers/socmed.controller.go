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

// CreateSocmed 	godoc
// @Summary 			create social media
// @Description 	create social media user
// @Tags 					social media
// @Accept 				application/json
// @Produce 			application/json
// @Param 				models.Socmed body models.Socmed true "socmed user"
// @Success 			200 {object} models.Socmed
// @Router 				/socmed/ [post]
func CreateSocmed(c *gin.Context) {
	db := databases.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socmed := models.Socmed{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = userID

	err := db.Debug().Create(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, Socmed)
}

// GetAllSocmed 	godoc
// @Summary 			get all social media
// @Description 	get all social media users
// @Tags 					social media
// @Accept 				application/json
// @Produce 			application/json
// @Success 			200 {object} models.Socmed
// @Router 				/socmed/ [get]
func GetAllSocmed(c *gin.Context) {
	db := databases.GetDB()

	Socmed := []models.Socmed{}

	err := db.Find(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"socmeds": Socmed,
	})
}

// GetSocmedById 	godoc
// @Summary 			get social media by id
// @Description 	get social media user by id
// @Tags 					social media
// @Accept 				application/json
// @Produce 			application/json
// @Param					socmedId path int true "id of the socmed"
// @Success 			200 {object} models.Socmed
// @Router 				/socmed/{socmedId} [get]
func GetSocmedById(c *gin.Context) {
	db := databases.GetDB()

	Socmed := models.Socmed{}
	socmedId, _ := strconv.Atoi(c.Param("socmedId"))

	err := db.First(&Socmed, "id = ?", socmedId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Socmed)
}

// UpdateSocmed 	godoc
// @Summary 			update social media
// @Description 	update social media user
// @Tags 					social media
// @Accept 				application/json
// @Produce 			application/json
// @Param					socmedId path int true "id of the socmed"
// @Success 			200 {object} models.Socmed
// @Router 				/socmed/{socmedId} [put]
func UpdateSocmed(c *gin.Context) {
	db := databases.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socmed := models.Socmed{}
	socmedId, _ := strconv.Atoi(c.Param("socmedId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = userID
	Socmed.ID = uint(socmedId)

	err := db.Model(&Socmed).Where("id = ?", socmedId).Updates(models.Socmed{
		Name:       Socmed.Name,
		Socmed_Url: Socmed.Socmed_Url,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Socmed)
}

// DeleteSocmed 	godoc
// @Summary 			delete social media
// @Description 	delete social media user
// @Tags 					social media
// @Accept 				application/json
// @Produce 			application/json
// @Param					socmedId path int true "id of the socmed"
// @Success 			200 {object} models.Socmed
// @Router 				/socmed/{socmedId} [delete]
func DeleteSocmed(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	Socmed := models.Socmed{}
	socmedId, _ := strconv.Atoi(c.Param("socmedId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	err := db.Where("id = ?", socmedId).Delete(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Socmed)
}
