package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func activeCmd() *cobra.Command {
	activeCMD := &cobra.Command{
		Use:     "active",
		Short:   "Active mode",
		Long:    `Avtive mode`,
		Example: `regoftw active`,
		Run:     runActive,
	}

	return activeCMD
}

func runActive(cmd *cobra.Command, args []string) {
	//Metadata Example
	// var functionsToExecute utils.ExecuteRegoFunction
	// If we have a good ctx, we probably do not need to pass parameters
	for _, d := range domainsToCheck {
		fmt.Println(d)
		// function := utils.RegoFunction{Function: functions.ExtractMetadata, Args: []string{d}}
		// functionsToExecute.Functions = append(functionsToExecute.Functions, function)
		// functionsToExecute.Run()
	}

}

func init() {
	rootCmd.AddCommand(activeCmd())
}
