package config

import (
	"fmt"

	cfgk "core/internal/config/key"
	"core/models"

	v "github.com/spf13/viper"
)

// GetRTApiToken Get formated RT api token
func GetRTApiToken() string {
	return fmt.Sprintf("token %s", v.GetString(cfgk.RTApiToken))
}

// Export RT Database Variables
func GetRTDatabaseHost() string {
	return v.GetString(cfgk.RTDatabaseHost)
}
func GetRTDatabasePort() string {
	return v.GetString(cfgk.RTDatabasePort)
}
func GetRTDatabaseName() string {
	return v.GetString(cfgk.RTDatabaseName)
}
func GetRTDatabaseUsername() string {
	return v.GetString(cfgk.RTDatabaseUsername)
}
func GetRTDatabasePassword() string {
	return v.GetString(cfgk.RTDatabasePassword)
}

// Export CSV Columns
func GetCSVColumns() *models.CSVColumns {
	return &models.CSVColumns{
		ID:                    int8(v.GetInt(cfgk.CSVColumnID)),
		Status:                int8(v.GetInt(cfgk.CSVColumnStatus)),
		TargetTeam:            int8(v.GetInt(cfgk.CSVColumnTargetTeam)),
		RootCause:             int8(v.GetInt(cfgk.CSVColumnRootCause)),
		ResolutionDescription: int8(v.GetInt(cfgk.CSVColumnResolutionDescription)),
		ONURxPower:            int8(v.GetInt(cfgk.CSVColumnRxPower)),
		Comment:               int8(v.GetInt(cfgk.CSVColumnComment)),
		SuspectedAreaOfIssue:  int8(v.GetInt(cfgk.CSVColumnSuspectedAreaOfIssue)),
		RootCauseCategory:     int8(v.GetInt(cfgk.CSVColumnRootCauseCategory)),
		ServiceRootCause:      int8(v.GetInt(cfgk.CSVColumnServiceRootCause)),
	}
}
