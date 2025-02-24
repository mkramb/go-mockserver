package models

type ApiResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}
