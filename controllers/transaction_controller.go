package controllers

import (
	interfaces "expensetracker/interfaces"
	"expensetracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	tm interfaces.TransactionManager
}

func NewTransactionController(tmi interfaces.TransactionManager) *TransactionController {
	return &TransactionController{tm: tmi}
}

// Add Transaction
func (tc TransactionController) AddTransaction(c *gin.Context) {
	params := tc.tm.Add()
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "addtransaction.tmpl", params)
}

// Save Transaction
func (tc TransactionController) SaveTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		return
	}

	params := tc.tm.Save(transaction)
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}

// Get Transactions
func (tc TransactionController) GetTransactions(c *gin.Context) {
	params := tc.tm.Get()
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "viewtransactions.tmpl", params)
}

// Edit Transaction
func (tc TransactionController) EditTransaction(c *gin.Context) {
	var params models.IdBinding
	if err := c.ShouldBind(&params); err != nil {
		return
	}
	c.JSON(http.StatusOK, params.Id)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}

// Update Transaction
func (tc TransactionController) UpdateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		return
	}

	params := tc.tm.Update(c.Query("id"), transaction)
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}

// Delete Transaction
func (tc TransactionController) DeleteTransaction(c *gin.Context) {
	var params models.IdBinding
	if err := c.ShouldBind(&params); err != nil {
		return
	}
	c.JSON(http.StatusOK, params.Id)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}
