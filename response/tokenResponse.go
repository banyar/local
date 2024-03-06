package response

type TokenErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
