package response

// SuccessResponse represents a successful API response.
type SuccessResponse[T any] struct {
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}