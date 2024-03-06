package key

const (
	CloudShareBaseUrl                     = "cloudshare.base_url"
	CloudSharePassword                    = "cloudshare.user.password"
	CloudShareUserName                    = "cloudshare.user.name"
	FileInputEnabled                      = "file.input.enabled"
	FileInputPathTemplate                 = "file.input.path"          // string (templated): file input path
	FileOutputDumpObjects                 = "file.output.dump_objects" // bool: Whether to dump memory objects to file
	FileOutputEnabled                     = "file.output.enabled"      // bool: Enable/Disable file output
	FileOutputPathTemplate                = "file.output.path"         // string (templated): File output path
	LoggingFileEnabled                    = "logging.file.enabled"
	LoggingFileNameTemplate               = "logging.file.name"
	LoggingFileTimestampInName            = "logging.file.timestamp_in_name"
	LoggingLevel                          = "logging.level"
	RTApiBaseUrl                          = "rt.api.base_url"
	RTApiToken                            = "rt.api.token"
	TicketUpdateCSVFileAge                = "rt.ticket.update.csv_file_age"       // duration: CSV file create/updated time diff from now (e.g., 1h, 30m, 3600s, etc.)
	TicketUpdateCheckStatusFirst          = "rt.ticket.update.check_status_first" // Check ticket status before proceeding to update
	TicketUpdateCloudShareFilePath_Input  = "rt.ticket.update.cloudshare.file_path.input"
	TicketUpdateCloudShareFilePath_Output = "rt.ticket.update.cloudshare.file_path.output"
	TicketUpdateUploadOutput2CloudShare   = "rt.ticket.update.upload_output_to_cloudshare"
	TicketUpdateUseCSVFromCloudShare      = "rt.ticket.update.use_csv_from_cloudshare" // bool: Whether to use csv file from CloudShare
)

// RT Database
const (
	RTDatabaseName     = "rt.database.name"
	RTDatabaseHost     = "rt.database.host"
	RTDatabasePort     = "rt.database.port"
	RTDatabaseUsername = "rt.database.username"
	RTDatabasePassword = "rt.database.password"
)

// CSV Column Indexes
const (
	CSVColumnID                    = "csv.columns.id"
	CSVColumnStatus                = "csv.columns.status"
	CSVColumnTargetTeam            = "csv.columns.target_team"
	CSVColumnRootCause             = "csv.columns.root_cause"
	CSVColumnResolutionDescription = "csv.columns.resolution_description"
	CSVColumnRxPower               = "csv.columns.rx_power"
	CSVColumnComment               = "csv.columns.comment"
	CSVColumnSuspectedAreaOfIssue  = "csv.columns.suspected_area_of_issue"
	CSVColumnRootCauseCategory     = "csv.columns.root_cause_category"
	CSVColumnServiceRootCause      = "csv.columns.service_root_cause"
)
