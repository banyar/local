package internal

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	cfgk "core/internal/config/key"
)

var (
	ReadInConfig bool
)

// initLogging sets logger settings
func InitLogging() {
	// Parse logging level
	loggingLevel, err := zapcore.ParseLevel(viper.GetString(cfgk.LoggingLevel))
	if err != nil {
		panic("error parsing logging level: " + err.Error())
	}
	// Multi-syncer with optional file syncer
	var syncers []zapcore.WriteSyncer
	syncers = append(syncers, zapcore.Lock(os.Stdout))
	// Determine the log file name using a dynamic template
	logFileNameTemplate := viper.GetString(cfgk.LoggingFileNameTemplate)
	includeTimestampInFileName := viper.GetBool(cfgk.LoggingFileTimestampInName)
	tmpl, err := template.New("logFileName").Parse(logFileNameTemplate)
	if err != nil {
		panic(err)
	}
	// Define data for the template
	data := struct {
		IncludeTimestamp bool
		Timestamp        time.Time
		WorkingDir       string
	}{
		IncludeTimestamp: includeTimestampInFileName,
		Timestamp:        time.Now(),
		WorkingDir:       WorkingDir,
	}
	// Create a buffer to capture the template execution result
	var logFileNameBuffer strings.Builder
	// Execute the template and write the result to the buffer
	if err = tmpl.Execute(&logFileNameBuffer, data); err != nil {
		panic(err)
	}
	logFileName := logFileNameBuffer.String() // Get the templated log file name from the buffer
	LoggingFileName = logFileName             // Re-set log file name with rendered name
	logFileDir := filepath.Dir(logFileName)   // Extract the directory path from the log file name
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(logFileDir, 0755); err != nil {
		fmt.Println("Error creating log directory:", err)
		return
	}
	// Setup file syncer
	fileSyncerEnabled := viper.GetBool(cfgk.LoggingFileEnabled)
	if fileSyncerEnabled {
		// If no user-provided log file name, generate file name based on whether to include timestamp
		if LoggingFileName != "" {
			logFileName = LoggingFileName
		}
		// Prepare file
		file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		// Add file syncer
		fileSyncer := zapcore.AddSync(file)
		syncers = append(syncers, fileSyncer)
	}
	// Create a multi-write syncer with all specified syncers
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(syncers...)
	// Create a new encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	// Use RFC3339 time format
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	// Create a console encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	// Create a core that uses the console encoder and the multi-write syncer
	core := zapcore.NewCore(consoleEncoder, multiWriteSyncer, loggingLevel)
	// Create a new logger with the core
	logger := zap.New(core)
	// Flushes any buffered log entries before the program exits.
	defer logger.Sync()
	// Set Logger in internal package
	Logger = logger
	// Write initial notifications to log
	logger.Debug("Log level", zap.String("level", logger.Level().String())) // Logging level
	// Log file name, if enabled
	if fileSyncerEnabled {
		// Get the absolute path
		logFileNameAbsPath, err := filepath.Abs(logFileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		logger.Debug("Log file", zap.String("name", logFileNameAbsPath))
	}
	// Config file
	if ReadInConfig {
		logger.Debug("Config file", zap.String("name", viper.ConfigFileUsed()))
	}
}
