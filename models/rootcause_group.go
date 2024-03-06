package models

type RootCauseGroup struct {
	SuspectedAreaOfIssue *map[string][]string
	RootCauseCategory    *map[string][]string
	ServiceRootCause     *map[string][]string
}
