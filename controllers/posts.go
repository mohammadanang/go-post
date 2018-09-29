package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-post/models"
)

func GetPosts(ctx *gin.Context) {
	var posts []models.PostItem
	// Error Handling
	if err := models.DB.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": posts,
	})
}

func CreatePost(ctx *gin.Context) {
	post := models.PostItem{Title: ctx.PostForm("title")}
	if err := models.DB.Save(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Create Failed",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": post,
		"message": "Post Created!",
		"status": http.StatusCreated,
	})
}

func ShowPost(ctx *gin.Context) {
	postId := ctx.Param("postId")
	var post models.PostItem
	// Error Handling
	if err := models.DB.First(&post, postId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	postId := ctx.Param("postId")
	var post models.PostItem
	// Error handling
	if err := models.DB.First(&post, postId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Data Not Found",
		})
		return
	}

	post.Title = ctx.PostForm("title")
	if err := models.DB.Save(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Update Failed!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": post,
	})
}

func DeletePost(ctx *gin.Context) {
	postId := ctx.Param("postId")
	var post models.PostItem
	// Error handling
	if err := models.DB.First(&post, postId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Data Not Found",
		})
		return
	}

	if err := models.DB.Delete(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Delete Failed!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Data Deleted!",
	})
}