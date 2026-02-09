package services

import (
	"gokasir-api/models"
	"gokasir-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItems) (*models.Transaction, error){
	return s.repo.CreateTransaction(items)
}