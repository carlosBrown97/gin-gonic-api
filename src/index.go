package src

import (
	"gin-gonic-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	newUser := controllers.NewUser()

	user := r.Group("/users")
	{
		user.GET("/", controllers.UsersGet(newUser))
		user.POST("/", controllers.UsersPost(newUser))
	}

	return r
}
