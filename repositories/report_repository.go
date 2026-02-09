package repositories

import (
	"database/sql"
	"gokasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) DailyReport() (*models.DailyReportResponse, error) {
	today := time.Now().Format("2006-01-02")
	// Get Transaction Count
	query := "SELECT id, total_amount FROM transactions WHERE created_at::date = $1"
	rows, err := r.db.Query(query, today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	totalAmount := 0
	totalTranction := 0
	var id, total, totTrx int
	for rows.Next() {
		err = rows.Scan(&id, &total)
		if err != nil {
			return nil, err
		}
		totalAmount += total

		query := "SELECT SUM(quantity) as quantity FROM transaction_details WHERE transaction_id = $1"
		err = r.db.QueryRow(query, id).Scan(&totTrx)
		if err != nil {
			return nil, err
		}
		totalTranction += totTrx
	}

	var bestSeller *models.BestSeller

	queryBestSeller := `SELECT p.name, SUM(td.quantity) AS qty_terjual 
						FROM transaction_details td
						JOIN products p ON p.id = td.product_id 
						JOIN transactions tx ON tx.id = td.transaction_id
						WHERE tx.created_at::date = $1
						GROUP BY p.id, p.name
						ORDER BY qty_terjual DESC
						LIMIT 1`

	var name string
	var qty int
	err = r.db.QueryRow(queryBestSeller, today).Scan(&name, &qty)
	if err == nil {
		bestSeller = &models.BestSeller{
			Name:         name,
			QuantitySell: qty,
		}
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	return &models.DailyReportResponse{
		TotalRevenue:     totalAmount,
		TotalTransaction: totalTranction,
		ProductTerlaris:  bestSeller,
	}, nil
}
