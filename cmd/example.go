/*
Copyright © 2023 Aung Zay Yar Lwin <aung.z.y.lwin@frontiir.net>
*/
package cmd

import (
	internal "core/internal"

	"errors"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	// "go.uber.org/zap"
)

// exampleCmd represents the example command
var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "Example sub command",
	Long:  `Shows example/demo results. For example: Sample log outputs with various levels are shown under this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		printSampleLogs()
	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printSampleLogs() {

	// Example logging statements:
	logger := internal.Logger
	logger.Info("This is an info message", zap.String("key", "value"))
	logger.Warn("This is a warning message", zap.Int("count", 42))
	logger.Error("This is an error message", zap.Error(errors.New("test error message")))

	// You can also use SugaredLogger for structured logging with placeholders.
	// sugar := logger.Sugar()
	// sugar.Infof("This is a formatted info message with placeholders: %s", "placeholder_value")

	// services.ValidateRootCause()
	// logger.Info("DB response", zap.("key", customFields))
	// logger.Error("Error", zap.Error(errors.New(err.Error.Error())))

}
