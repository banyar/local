/*
Copyright © 2023 Aung Zay Yar Lwin <aung.z.y.lwin@frontiir.net>
*/
package config

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config related sub-commands",
	Long: `Config related sub-commands

Run one of the available sub commands.`,
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Export configCmd variable
func GetConfigCmd() *cobra.Command {
	return configCmd
}
