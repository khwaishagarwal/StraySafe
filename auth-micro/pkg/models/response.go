package models

type JsonResponse struct {
	Message string `json:"message"`
	Content any    `json:"content,omitempty"`
}
