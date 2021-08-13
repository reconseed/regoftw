package cmd

import (
	"fmt"
	"regoftw/functions"
	"regoftw/utils"

	"github.com/spf13/cobra"
)

func reconCmd() *cobra.Command {
	reconCMD := &cobra.Command{
		Use:     "recon",
		Short:   "Recon mode",
		Long:    `Recon mode`,
		Example: `regoftw recon`,
		Run:     runReconMode,
	}

	return reconCMD
}

func runReconMode(cmd *cobra.Command, args []string) {
	//Metadata Example
	var functionsToExecute utils.ExecuteRegoFunction
	// If we have a good ctx, we probably do not need to pass parameters
	for _, d := range domainsToCheck {
		fmt.Println(d)
		function := utils.RegoFunction{Function: functions.ExtractMetadata, Args: []string{d}}
		functionsToExecute.Functions = append(functionsToExecute.Functions, function)
		functionsToExecute.Run()
	}

}

func init() {
	rootCmd.AddCommand(reconCmd())
}
