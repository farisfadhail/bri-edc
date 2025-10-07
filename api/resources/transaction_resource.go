package resources

type SaleResource struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

type SettlementResource struct {
	BatchID     string `json:"batch_id"`
	TotalCount  int    `json:"total_count"`
	Approved    int    `json:"approved"`
	Declined    int    `json:"declined"`
	TotalAmount int    `json:"total_amount"`
}
