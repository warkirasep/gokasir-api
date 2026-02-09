package handlers

import (
	"encoding/json"
	"gokasir-api/services"
	"net/http"
)

type ReportHandler struct {
	services *services.ReportService
}

func NewReportHandler(services *services.ReportService) *ReportHandler {
	return &ReportHandler{services: services}
}

func (h *ReportHandler) DailyReport(w http.ResponseWriter, r *http.Request){

	report, err := h.services.DailyReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) HandleDailyReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.DailyReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}