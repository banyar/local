package models

import "github.com/SamuelTissot/sqltime"

type Ticket struct {
	ID              int32        `json:"id" gorm:"primaryKey"`
	EffectiveId     int32        `json:"effective_id"`
	Queue           int32        `json:"queue"`     // Queue table id
	IsMerged        int32        `json:"is_merged"` //
	Type            string       `json:"type"`
	Owner           int32        `json:"owner"`
	Subject         string       `json:"subject"`
	InitialPriority int32        `json:"initial_priority"`
	FinalPriority   int32        `json:"final_priority"`
	Priority        int32        `json:"priority"`
	TimeEstimated   int32        `json:"time_estimated"`
	TimeWorked      int32        `json:"time_worked"`
	Status          string       `json:"status"`
	SLA             string       `json:"sla"`
	TimeLeft        int32        `json:"timeleft"`
	Told            sqltime.Time `json:"told"`
	Starts          sqltime.Time `json:"starts"`
	Started         sqltime.Time `json:"started"`
	Due             sqltime.Time `json:"due"`
	Resolved        sqltime.Time `json:"resolved"`
	LastUpdatedBy   int32        `json:"last_updated_by"`
	LastUpdated     sqltime.Time `json:"last_updated"`
	Creator         int32        `json:"creator"`
	Created         sqltime.Time `json:"created"`
}

func (c *Ticket) TableName() string {
	// ORM mapping table name, this is default
	return "Tickets"
}
