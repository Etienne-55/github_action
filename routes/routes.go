package routes

import "github.com/gin-gonic/gin"


func SetupRoutes(r *gin.Engine) {


	api := r.Group("/api")
	{
		api.GET("hello",)
	}


	r.StaticFile("/hello", "./templates/hello.html")


}
