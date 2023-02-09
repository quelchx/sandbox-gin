package main

import (
	"github.com/quelchx/go-crud/initializers"
	model "github.com/quelchx/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.DatabaseClient()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
