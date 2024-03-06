package main

import (
	"core/cmd"
	"core/utils"
)

func main() {

	cmd.Execute()

	utils.DisplayJsonFormat("", "====================================================")
	// result, _ := services.GetTicketByID(uint(3924319))
	// utils.DisplayJsonFormat("TEST", result)
	// data := services.GetFormattedRootCauseValues()

	// utils.DisplayJsonFormat("MAIN -----------> ", data)

}
