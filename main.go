package main

import (
	"expensetracker/data"
	"expensetracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	data.TryConnect()

	routes.ConfigureWebRoutes(r)
	r.Run(":8080")
}
