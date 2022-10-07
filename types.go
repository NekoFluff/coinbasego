package coinbase

type ProductID string

type PaginationParams struct {
	Before string `url:"before,omitempty"`
	After  string `url:"after,omitempty"`
	Limit  int    `url:"limit" binding:"required"`
}

type PaginationResponse struct {
	Before string
	After  string
}
