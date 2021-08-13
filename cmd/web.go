package cmd

import (
	"github.com/spf13/cobra"
)

func webCmd() *cobra.Command {
	webCMD := &cobra.Command{
		Use:     "web",
		Short:   "Web mode",
		Long:    `Web mode`,
		Example: `regoftw web`,
		Run:     runWebMode,
	}

	return webCMD
}

func runWebMode(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(webCmd())
}
