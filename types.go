package coinbasego

type ProductID string

type PaginationParams struct {
	Cursor string `url:"cursor,omitempty"`
	Limit  int    `url:"limit" binding:"required"`
}

type PaginationResponse struct {
	HasNext bool   `json:"has_next"`
	Cursor  string `json:"cursor"`
	Size    int32  `json:"size"`
}
