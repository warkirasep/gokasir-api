package models

type DailyReportResponse struct {
	TotalRevenue     int         `json:"total_revenue"`
	TotalTransaction int         `json:"total_transaksi"`
	ProductTerlaris  *BestSeller `json:"produk_terlaris"`
}

type BestSeller struct {
	Name         string `json:"nama"`
	QuantitySell int    `json:"qty_terjual"`
}
