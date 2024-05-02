package models

type JsonResponse struct {
	Message string `json:"message,omitempty"`
	Content any    `json:"content,omitempty"`
}
