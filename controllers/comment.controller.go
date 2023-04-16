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

// CreateComment 	godoc
// @Summary 			create comment
// @Description 	create comment user
// @Tags 					comment
// @Accept 				application/json
// @Produce 			application/json
// @Param 				models.Comment body models.Comment true "comment user"
// @Success 			200 {object} models.Comment
// @Router 				/comment/ [post]
func CreateComment(c *gin.Context) {
	db := databases.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// GetAllComment 	godoc
// @Summary 			get all comment
// @Description 	get all comment users
// @Tags 					comment
// @Accept 				application/json
// @Produce 			application/json
// @Success 			200 {object} models.Comment
// @Router 				/comment/ [get]
func GetAllComment(c *gin.Context) {
	db := databases.GetDB()

	Comment := []models.Comment{}

	err := db.Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": Comment,
	})
}

// GetCommentById 	godoc
// @Summary 			get comment by id
// @Description 	get comment user by id
// @Tags 					comment
// @Accept 				application/json
// @Produce 			application/json
// @Param					commentId path int true "id of the comment"
// @Success 			200 {object} models.Comment
// @Router 				/comment/{commentId} [get]
func GetCommentById(c *gin.Context) {
	db := databases.GetDB()

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// UpdateComment 	godoc
// @Summary 			update comment
// @Description 	update comment user
// @Tags 					comment
// @Accept 				application/json
// @Produce 			application/json
// @Param					commentId path int true "id of the comment"
// @Success 			200 {object} models.Comment
// @Router 				/comment/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := databases.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comment.Message,
		PhotoID: Comment.PhotoID,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment 	godoc
// @Summary 			delete comment
// @Description 	delete comment user
// @Tags 					comment
// @Accept 				application/json
// @Produce 			application/json
// @Param					commentId path int true "id of the comment"
// @Success 			200 {object} models.Comment
// @Router 				/comment/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Comment)
}
