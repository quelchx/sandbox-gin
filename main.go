package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quelchx/go-crud/initializers"
	"github.com/quelchx/go-crud/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.DatabaseClient()
}

func main() {
	r := gin.Default()

	// use the post routes
	routers.PostRoutes(r)

	r.Run() // listen and serve on
}
