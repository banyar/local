/*
Copyright © 2023 Aung Zay Yar Lwin <aung.z.y.lwin@frontiir.net>
*/
package ticket

import (
	"github.com/spf13/cobra"
)

// ticketCmd represents the ticket command
var ticketCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Ticket related sub-commands",
	Long: `Ticket related sub-commands.

Run one of the available sub commands.`,
}

func init() {
	// rootCmd.AddCommand(ticketCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ticketCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ticketCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Export ticketCmd variable
func GetTicketCmd() *cobra.Command {
	return ticketCmd
}
