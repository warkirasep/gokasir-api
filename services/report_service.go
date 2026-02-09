package services

import (
	"gokasir-api/models"
	"gokasir-api/repositories"
)

type ReportService struct{
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService{
	return &ReportService{repo: repo}
}

func (s *ReportService) DailyReport() (*models.DailyReportResponse, error){
	return s.repo.DailyReport()
}