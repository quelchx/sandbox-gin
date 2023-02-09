// router for post controller
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/quelchx/go-crud/controllers"
)

func PostRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/posts", controllers.GetPosts)
	incomingRoutes.POST("/posts", controllers.PostsCreate)
	incomingRoutes.GET("/posts/:id", controllers.GetPostById)
	incomingRoutes.PUT("/posts/:id", controllers.UpdatePostById)
	incomingRoutes.DELETE("/posts/:id", controllers.DeletePostById)
}
