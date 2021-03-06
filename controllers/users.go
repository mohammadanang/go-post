package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-post/models"
)

type FormUser struct {
	Username	string `form:"username"`
	Password	string `form:"password"`
}

type RawUser struct {
	Username	string `json:"username"`
	Password	string `json:"password"`
}

func Register(context *gin.Context) { // Test Post
	username := context.PostForm("username")
	context.String(http.StatusOK, "Hello lagi %s ", username)
}

func RegisterForm(context *gin.Context) {
	var user FormUser
	data := context.ShouldBind(&user)
	if data != nil {
		context.String(http.StatusInternalServerError, "invalid gaes")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Profile(context *gin.Context) { // Test Get
	username := context.Query("username")
	context.String(http.StatusOK, "Hello %s ", username)
}

func Category(context *gin.Context) { // Test from Path
	username := context.Param("username")
	context.String(http.StatusOK, "Hello %s ", username)
}

// From database
func ShowUser(ctx *gin.Context) {
	var user []models.User
	err := models.DB.Find(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{} {
			"status": http.StatusInternalServerError,
			"message": "Gagal",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string] interface{}{
		"status": http.StatusOK,
		"data": user,
	})
}

// From database with alias
func ShowPosting(ctx *gin.Context) {
	var posting []models.PostItem
	err := models.DB.Find(&posting).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status": http.StatusInternalServerError,
			"message": "Gagal",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string] interface{}{
		"status": http.StatusOK,
		"data": posting,
	})
}

// Function with where
func ShowDetailUser(ctx *gin.Context) {
	id := ctx.Query("username")
	var user models.User
	err := models.DB.Where("username = ?", id).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status": http.StatusInternalServerError,
			"message": "Gagal",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string] interface{}{
		"status": http.StatusOK,
		"data": user,
	})
}

// Function create
func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//err = user.ValidationUser()
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
	//		"message": err.Error(),
	//		"status": http.StatusInternalServerError,
	//	})
	//	return
	//}
	//
	//err = models.DB.Create(&user).Error
	//if err != nil {
	//	ctx.JSON(500, map[string] interface{}{
	//		"status": http.StatusInternalServerError,
	//		"message": "error",
	//	})
	//	return
	//}

	ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
		"status": http.StatusOK,
		"message": user,
	})
}

func MigrateTable(ctx *gin.Context) {
	err := models.DB.AutoMigrate(&models.Activities{}, &models.User{}, &models.PostItem{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status": http.StatusForbidden,
			"message": "Gagal dimigrasi",
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
		"status": http.StatusOK,
		"message": "Berhasil migrasi",
	})
}

// Create middleware
//func CheckHeaderAuthorization(ctx *gin.Context) {
//	authorization := ctx.GetHeader("Authorization")
//	if authorization != "12345" { // Sample with hardcode 12345
//		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
//			"message": "Unauthorized",
//		})
//		ctx.Abort{}
//	}
//	ctx.Next()
//}
