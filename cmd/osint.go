package cmd

import (
	"github.com/spf13/cobra"
)

func osintCmd() *cobra.Command {
	osintCMD := &cobra.Command{
		Use:     "osint",
		Short:   "Osint mode",
		Long:    `Osint mode`,
		Example: `regoftw osint`,
		Run:     runOsintMode,
	}

	return osintCMD
}

func runOsintMode(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(osintCmd())
}
