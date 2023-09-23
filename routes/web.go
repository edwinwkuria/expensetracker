package routes

import (
	"expensetracker/controllers"
	"expensetracker/services"

	"github.com/gin-gonic/gin"
)

func ConfigureWebRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("./templates/*")
	ts := services.NewTransactionManager()
	tc := controllers.NewTransactionController(ts)
	r.GET("/", tc.GetTransactions)
	r.GET("/:id", tc.EditTransaction)
	r.GET("/add", tc.AddTransaction)
	r.POST("/:id", tc.SaveTransaction)
	r.PATCH("/:id", tc.UpdateTransaction)
	r.DELETE("/:id", tc.DeleteTransaction)
}
