package handlers

import (
	"encoding/json"
	"gokasir-api/models"
	"gokasir-api/services"
	"net/http"
)

type TransactionHandler struct {
	services *services.TransactionService
}

func NewTransactionHandler(services *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{services: services}
}

func (h *TransactionHandler) Checkout (w http.ResponseWriter, r *http.Request){
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.services.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}