/*
Copyright © 2023 Aung Zay Yar Lwin <aung.z.y.lwin@frontiir.net>
*/
package ticket

import (
	"github.com/spf13/cobra"
)

var (
	ticketId int
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an RT ticket info",
	Long: `Get an RT ticket info.

Currently the cli will fetch ticket info via REST API.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTicketInfo(ticketId)
	},
}

func init() {
	ticketCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().IntVarP(&ticketId, "id", "i", 0, "Ticket ID - integer value of the ticket ID.")
	getCmd.MarkFlagRequired("id")
}

// getTicketInfo Get ticket info by calling Api
func getTicketInfo(ticketId int) {

}
