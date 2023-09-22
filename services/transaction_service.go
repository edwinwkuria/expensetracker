package services

import (
	"expensetracker/data"
	"expensetracker/models"
)

type TransactionService struct {
}

func NewTransactionManager() *TransactionService {
	return &TransactionService{}
}

func (ms *TransactionService) Add() bool {
	return true
}
func (ms *TransactionService) Save(t models.Transaction) bool {
	return data.SaveTransaction(t)
}
func (ms *TransactionService) Get() []models.Transaction {
	return data.GetTransactions()
}
func (ms *TransactionService) Edit(id string) models.Transaction {
	return data.GetTransaction(id)
}
func (ms *TransactionService) Update(id string, t models.Transaction) bool {
	return data.UpdateTransaction(id, t)
}
func (ms *TransactionService) Delete(id string) bool {
	return data.DeleteTransaction(id)
}
