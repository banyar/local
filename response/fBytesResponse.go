package response

// TicketResponse ...
type FBytesResponse struct {
	Status int           `json:"status"`
	Data   *FBytesDetail `json:"data"`
}

type FBytesDetail struct {
	ID           int32               `json:"id"`
	CustomFields *FBytesCustomFields `json:"custom_fields"`
}

type FBytesCustomFields struct {
	LocalServiceId   string         `json:"local_service_id"`
	CustomerId       string         `json:"customer_id"`
	CpeId            string         `json:"cpe_id"`
	TicketProblem    string         `json:"ticket_problem"`
	ServiceType      string         `json:"service_type"`
	InstallationType string         `json:"installation_type"`
	MainRootCause    *MainRootCause `json:"root_causes"`
}

type MainRootCause struct {
	SuscpectedAreaOfIssue string `json:"suspected_area_of_issue"`
	RootCauseCategory     string `json:"root_cause_category"`
	ServiceRootCause      string `json:"service_root_cause"`
}
