package cmd

import (
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

}

func init() {
	rootCmd.AddCommand(activeCmd())
}
