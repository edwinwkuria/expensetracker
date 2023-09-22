package models

type TransactionType int64

const (
	Debit TransactionType = iota
	Credit
)
