/*
Copyright © 2023 Aung Zay Yar Lwin <aung.z.y.lwin@frontiir.net>
*/
package cmd

import (
	"core/cmd/config"
	"core/cmd/ticket"
	internal "core/internal"

	viper "github.com/spf13/viper"

	cfgk "core/internal/config/key"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rtutil",
	Short: "Universal commandline utility for RT",
	Long: `This cli is a universal commandline utility to do various functions with 
Frontiir's RT ticket management system.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("Version", internal.Version())
	rootCmd.Version = internal.Version()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	internal.AppName = rootCmd.Use // Set the rootCmd's cmd name as AppName
	internal.AppStartTime = time.Now()
	// Initialize sub-commands
	rootCmd.AddCommand(config.GetConfigCmd())
	rootCmd.AddCommand(ticket.GetTicketCmd())

	// Initialize
	cobra.OnInitialize(initConfig, internal.InitLogging)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config",
		"", fmt.Sprintf("\"config file (default is $PWD/.%s.%s)\"", rootCmd.Use, internal.ConfigType))
	rootCmd.PersistentFlags().StringVar(&internal.LoggingFileName, internal.RootCPFlagLogFileName,
		"", "Log file name")
	rootCmd.PersistentFlags().StringVar(&internal.LoggingLevel, internal.RootCPFlagLogLevel,
		"", "Logging level (fatal, panic, dpanic, error, warn, info, debug)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Set viper config from flags
	viper.BindPFlag(cfgk.LoggingLevel, rootCmd.PersistentFlags().Lookup(internal.RootCPFlagLogLevel))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find working directory.
		wd, err := os.Getwd()
		cobra.CheckErr(err)
		internal.WorkingDir = wd // Set working dir var
		// Search config in working directory with name ".rtutil" (without extension).
		viper.AddConfigPath(wd)
		viper.SetConfigType(internal.ConfigType)
		viper.SetConfigName(fmt.Sprintf(".%s", rootCmd.Use))
		//viper.SetConfigName(".rtutil")
	}
	defaultFilePathTemplate := fmt.Sprintf(
		"{{.WorkingDir}}/.tmp/{{.Timestamp.Format \"%s\"}}", internal.DTFormatDirName) // I/O file path template
	// Set default config values
	viper.SetDefault(cfgk.CloudShareBaseUrl, internal.DefaultCloudShareBaseUrl)
	viper.SetDefault(cfgk.CloudSharePassword, "password")
	viper.SetDefault(cfgk.CloudShareUserName, "username")
	viper.SetDefault(cfgk.FileInputEnabled, true)
	viper.SetDefault(cfgk.FileInputPathTemplate, defaultFilePathTemplate)
	viper.SetDefault(cfgk.FileOutputDumpObjects, true)
	viper.SetDefault(cfgk.FileOutputEnabled, true)
	viper.SetDefault(cfgk.FileOutputPathTemplate, defaultFilePathTemplate)
	viper.SetDefault(cfgk.LoggingFileEnabled, true)
	viper.SetDefault(cfgk.LoggingFileNameTemplate, fmt.Sprintf(
		"%s/{{if .IncludeTimestamp}}{{.Timestamp.Format \"%s\"}}{{else}}%s{{end}}.log",
		defaultFilePathTemplate, internal.DTFormatFilename, rootCmd.Use,
	))
	viper.SetDefault(cfgk.LoggingFileTimestampInName, true)
	viper.SetDefault(cfgk.LoggingLevel, "info")
	viper.SetDefault(cfgk.RTApiBaseUrl, internal.DefaultRTApiBaseUrl)
	viper.SetDefault(cfgk.RTApiToken, "rt-api-token")
	viper.SetDefault(cfgk.RTDatabaseHost, "host")
	viper.SetDefault(cfgk.RTDatabasePort, "port")
	viper.SetDefault(cfgk.RTDatabaseName, "dbname")
	viper.SetDefault(cfgk.RTDatabaseUsername, "username")
	viper.SetDefault(cfgk.RTDatabasePassword, "password")
	viper.SetDefault(cfgk.TicketUpdateCSVFileAge, "15m")
	viper.SetDefault(cfgk.TicketUpdateCheckStatusFirst, true)
	viper.SetDefault(cfgk.TicketUpdateCloudShareFilePath_Input, internal.DefaultCloudShareFilePath_Input)
	viper.SetDefault(cfgk.TicketUpdateCloudShareFilePath_Output, internal.DefaultCloudShareFilePath_Output)
	viper.SetDefault(cfgk.TicketUpdateUploadOutput2CloudShare, true)
	viper.SetDefault(cfgk.TicketUpdateUseCSVFromCloudShare, true)

	// read in environment variables that match
	viper.AutomaticEnv()
	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		internal.ReadInConfig = true
	}

}
