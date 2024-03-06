package ticket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"core/internal/config"

	cfgk "core/internal/config/key"

	cfg "github.com/spf13/viper"
)

// Ticket object to get from RT
type Ticket struct {
	Id          int       `json:"id"`
	Created     time.Time `json:"Created"`
	LastUpdated time.Time `json:"LastUpdated"`
	Resolved    time.Time `json:"Resolved"`
	Started     time.Time `json:"Started"`
	Status      string    `json:"Status"`
	Subject     string    `json:"Subject"`
	Type        string    `json:"Type"`
}

// StatusPayload Payload for Ticket status POST/PUT payload
type StatusPayload struct {
	Status string `json:"Status"`
}

// GetTicket This function calls RT's REST API to get ticket info
func GetTicket(ticketId int) (*http.Response, error) {
	// logger := inl.Logger
	apiUrl := fmt.Sprintf("%s/ticket/%d", cfg.GetString(cfgk.RTApiBaseUrl), ticketId)
	// Create a new http.Client object
	client := &http.Client{}
	// Create a new http.Request object
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	// Set the Authorization header
	req.Header.Add("Authorization", config.GetRTApiToken())
	// Make the request
	resp, err := client.Do(req)
	// return
	return resp, err
}

// ChangeStatus Change Ticket's status
func ChangeStatus(ticketId string, status string) (*http.Response, error) {
	// HTTP api call
	apiUrl := fmt.Sprintf("%s/ticket/%s", cfg.GetString(cfgk.RTApiBaseUrl), ticketId)
	// create payload struct
	payloadObj := struct {
		Status string `json:"Status"`
	}{Status: status}
	// Serialize the data to JSON
	payload, err := json.Marshal(payloadObj)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON. %v", err)
	}
	client := &http.Client{} // Create a new http.Client object
	// Create a PUT request
	req, err := http.NewRequest("PUT", apiUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	req.Header.Add("Authorization", config.GetRTApiToken()) // Set the Authorization header
	req.Header.Set("Content-Type", "application/json")      // Set the Content-Type header to indicate JSON data
	// Make the PUT request
	resp, err := client.Do(req)
	return resp, err
}
