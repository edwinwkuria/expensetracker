package interfaces

import "expensetracker/models"

type TransactionManager interface {
	Add() bool
	Save(t models.Transaction) bool
	Get() []models.Transaction
	Edit(id string) models.Transaction
	Update(id string, t models.Transaction) bool
	Delete(id string) bool
}
