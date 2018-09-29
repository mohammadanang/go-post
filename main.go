package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-post/controllers"
)

func main() {
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"username" : "CryptoX",
		})
	})
	route.POST("/register", controllers.Register)
	route.POST("/register-form", controllers.RegisterForm)
	// Post routes
	postRoute := route.Group("/post")
	{
		postRoute.GET("/", controllers.GetPosts)
		postRoute.POST("/create", controllers.CreatePost)
		postRoute.GET("/show/:postId", controllers.ShowPost)
		postRoute.PUT("/update/:postId", controllers.UpdatePost)
		postRoute.DELETE("/delete/:postId", controllers.DeletePost)
	}
	// User routes
	userRoute := route.Group("/user")
	{
		userRoute.GET("/profile", controllers.Profile)
		userRoute.GET("/profile/:username", controllers.Category)
		userRoute.GET("/show", controllers.ShowUser)
		userRoute.GET("/show-post", controllers.ShowPosting)
		userRoute.GET("/detail", controllers.ShowDetailUser)
		userRoute.POST("/create", controllers.CreateUser)
	}
	// Migrate table
	route.GET("/migrate", controllers.MigrateTable)
	route.Run(":8888")
}
