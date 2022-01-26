package models

type GetPetsResponse []struct {
	ID    int     `json:"id"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type PostPetsRequest struct {
	ID    int     `json:"id" binding:"required"`
	Type  string  `json:"type" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type PostPetsResponse struct {
	Pet struct {
		ID    int     `json:"id"`
		Type  string  `json:"type"`
		Price float64 `json:"price"`
	} `json:"pet"`
	Message string `json:"message"`
}
