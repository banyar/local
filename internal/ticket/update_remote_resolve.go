package ticket

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

/*
This file includes code for updating tickets for remote-resolve results
*/

type CSVReadNote struct {
	FileName     string    `json:"fileName"`     // CSV file name
	Success      bool      `json:"success"`      // Status showing whether reading file is successgul
	AccessedTime time.Time `json:"accessedTime"` // File reading time
	Note         string    `json:"note"`         // Note
}

// CSVRecRemoteResolvedTicket CSV record for remote-resolved ticket
type CSVRecRemoteResolvedTicket struct {
	ID                    string `json:"id"`                    // RT ticket ID
	CurrentStatus         string `json:"currentStatus"`         // RT ticket current status
	TargetTeam            string `json:"targetTeam"`            // RT Queue
	RootCause             string `json:"rootCause"`             // RT CF: Root Cause
	ResolutionDescription string `json:"resolutionDescription"` // RT CF: Resolution Description
	ONURxPower            string `json:"onuRxPower"`            // RT CF: Resolved Fiber ONU Signal dBm
	Comment               string `json:"comment"`               // RT CF: Comment
	SuspectedAreaOfIssue  string `json:"suspectedAreaOfIssue"`  // RT CF: Suspected Area Of Issue
	RootCauseCategory     string `json:"rootCauseCategory"`     // RT CF: Root Cause Category
	ServiceRootCause      string `json:"serviceRootCause"`      // RT CF: Service Root Cause
}

// CustomFieldsRemoteResolved Ticket Custom Fields for Remote-resolved update
type CustomFieldsRemoteResolved struct {
	RootCause             string `json:"Root Cause"`
	ResolutionDescription string `json:"Resolution Description"`
	ONURxPower            string `json:"Resolved Fiber ONU Signal dBm"`
	SuspectedAreaOfIssue  string `json:"Suspected Area Of Issue"` // RT CF: Suspected Area Of Issue
	RootCauseCategory     string `json:"Root Cause Category"`     // RT CF: Root Cause Category
	ServiceRootCause      string `json:"Service Root Cause"`      // RT CF: Service Root Cause
}

// UpdateTicketRemoteResolved Ticket record for Remote-resolved update
type UpdateTicketRemoteResolved struct {
	Id           int                        `json:"id"` // RT ticket ID
	Status       string                     `json:"Status"`
	Queue        string                     `json:"Queue"`
	CustomFields CustomFieldsRemoteResolved `json:"CustomFields"`
}

type Rest2PayloadComments struct {
	Content     string `json:"Content"`
	Subject     string `json:"Subject"`
	ContentType string `json:"ContentType"`
}

// UpdateTicketRemoteResolvedLog Log entry for ticket update remote-resolved
type UpdateTicketRemoteResolvedLog struct {
	Request  UpdateTicketRemoteResolved
	Response UpdateTicketRemoteResolvedLogResponse
}

// UpdateTicketRemoteResolvedLogResponse Response body for update log
type UpdateTicketRemoteResolvedLogResponse struct {
	Status     string
	StatusCode int
	Body       interface{}
}

// GetONURxPowerRoundedNumString Get 2 decimal places rounded number string.
//
// The method returns "" if t.ONURxPower == "".
func (t *CSVRecRemoteResolvedTicket) GetONURxPowerRoundedNumString() string {
	if t.ONURxPower == "" {
		return t.ONURxPower
	} else {
		f, _ := strconv.ParseFloat(t.ONURxPower, 64) // Parse string to float
		//fmt.Printf("parsed number: %v\n", f)
		f = math.Round(f*100) / 100 // multiply with 100, round number, divided by 100
		//fmt.Printf("rouended number: %v\n", f)
		s := strconv.FormatFloat(f, 'f', 2, 64)
		//fmt.Printf("string: %v\n", s)
		return s
	}
}

// GetIDInt Integer ID value for CSVRecRemoteResolvedTicket
func (t *CSVRecRemoteResolvedTicket) GetIDInt() int {
	id, _ := strconv.Atoi(t.ID)
	return id
}

// SetONURxPowerString Convert ONURxPower from float to string
func (c *CustomFieldsRemoteResolved) SetONURxPowerString(onuRxPower float32) {
	c.ONURxPower = fmt.Sprintf("%f", onuRxPower)
}
