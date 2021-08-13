package cmd

import (
	"github.com/spf13/cobra"
)

func subdomainsCmd() *cobra.Command {
	subdomainsCMD := &cobra.Command{
		Use:     "subdomains",
		Short:   "Subdomains mode",
		Long:    `Subdomains mode`,
		Example: `regoftw subdomains`,
		Run:     runSubdomainsMode,
	}

	return subdomainsCMD
}

func runSubdomainsMode(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(subdomainsCmd())
}
