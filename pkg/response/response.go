package response

// SuccessResponse represents a successful API response.
type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}