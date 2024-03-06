package internal

import (
	"time"

	"go.uber.org/zap"
)

const (
	// DateTime Format Strings
	DTFormatDirName  = "2006-01-02"
	DTFormatFilename = "2006-01-02_1504"

	// Misc
	ConfigType                       = "json"
	DefaultCloudShareBaseUrl         = "https://cloudshare.frontiir.net/remote.php/webdav"
	DefaultCloudShareFilePath_Input  = "/Shared Operations/RnD/Operations Support/RT/Ticket/UpdateRemoteResolved/CSV_Upload"
	DefaultCloudShareFilePath_Output = "/Shared Operations/RnD/Operations Support/RT/Ticket/UpdateRemoteResolved/UpdateJobOutputs"
	DefaultRTApiBaseUrl              = "https://rt-dev.frontiir.net/REST/2.0"

	// Ticket status
	TktStatusCancelled  = "Cancelled"
	TktStatusClosed     = "Closed"
	TktStatusInProgress = "In_Progress"
	TktStatusNew        = "New"
	TktStatusPending    = "Pending"
	TktStatusReOpen     = "Re-Open"
	TktStatusReOpen2    = "Re-Open-2"
	TktStatusReOpen3    = "Re-Open-3"
	TktStatusReOpen4    = "Re-Open-4"
	TktStatusResolved   = "Resolved"
	TktStatusResolved2  = "Resolved-2"
	TktStatusResolved3  = "Resolved-3"
	TktStatusResolved4  = "Resolved-4"
	TktStatusResolved5  = "Resolved-5"
)

var (
	AppEndTime                 time.Time
	AppName                    string
	AppStartTime               time.Time
	FilePath_Input             string
	FilePath_Output            string
	Logger                     *zap.Logger
	LoggingFileName            string
	LoggingFileTimestampInName bool
	LoggingLevel               string
	WorkingDir                 string

	// Command [Persistent flags]
	// Prefix meanings:
	//   CPFlag - Cmd Persistent Flag
	//   CFlag - Cmd Flag
	CPFlagUseCSVFromCoudShare = "cloudshare"
	RootCPFlagLogFileName     = "log-file"
	RootCPFlagLogLevel        = "log-level"
)
