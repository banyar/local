package response

// TicketResponse ...
type TicketResponse struct {
	Status int           `json:"status"`
	Data   *TicketDetail `json:"data"`
}

type TicketDetail struct {
	ID           int32             `json:"id"`
	CustomFields map[string]string `json:"custom_fields"`
}
