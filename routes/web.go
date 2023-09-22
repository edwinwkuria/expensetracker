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
	r.GET("/edit", tc.EditTransaction)
	r.GET("/add", tc.AddTransaction)
	r.POST("/save", tc.SaveTransaction)
	r.PATCH("/update", tc.UpdateTransaction)
	r.DELETE("/delete", tc.DeleteTransaction)
}
