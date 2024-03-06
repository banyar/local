/*
Copyright © 2023 Aung Zay Yar Lwin <aung.z.y.lwin@frontiir.net>
*/
package config

import (
	"github.com/spf13/cobra"
)

var (
	cfgFileName  = ""
	overrideFile = false // Override existing file
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate config file",
	Long: `Generate config file with default values.

This sub command generates a config file with default values.

Unless used --force flag, new config file will not be generated if a file with same name exists in current directory.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	configCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringVarP(&cfgFileName, "filename", "f", cfgFileName, "Custom config file name")
	generateCmd.Flags().BoolVarP(&overrideFile, "force", "", false, "Force generate config file. If true, the existing file with same name will be overriden")
}
