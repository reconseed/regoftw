package cmd

import (
	"github.com/spf13/cobra"
)

func fullCmd() *cobra.Command {
	fullCMD := &cobra.Command{
		Use:     "full",
		Short:   "Full mode",
		Long:    `Full mode`,
		Example: `regoftw full`,
		Run:     runFull,
	}

	return fullCMD
}

func runFull(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(fullCmd())
}
