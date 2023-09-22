package controllers

import (
	interfaces "expensetracker/interfaces"
	"expensetracker/models"
	"net/http"
	"strconv"

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
	intValue, err := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	if err != nil {
		return
	}

	transaction := models.Transaction{
		Reference:       c.PostForm("ref"),
		Description:     c.PostForm("desc"),
		TransactionType: c.PostForm("trantype"),
		Amount:          float64(intValue),
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

	params := tc.tm.Edit(c.Query("id"))
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}

// Update Transaction
func (tc TransactionController) UpdateTransaction(c *gin.Context) {
	intValue, err := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	if err != nil {
		return
	}
	transaction := models.Transaction{
		Reference:       c.PostForm("ref"),
		Description:     c.PostForm("desc"),
		TransactionType: c.PostForm("trantype"),
		Amount:          float64(intValue),
	}

	params := tc.tm.Update(c.Query("id"), transaction)
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}

// Delete Transaction
func (tc TransactionController) DeleteTransaction(c *gin.Context) {
	params := tc.tm.Delete(c.Query("id"))
	c.JSON(http.StatusOK, params)
	//c.HTML(http.StatusOK, "viewtransaction.tmpl", params)
}
