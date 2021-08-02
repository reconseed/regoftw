package cmd

import (
	"regoftw/utils"

	"github.com/spf13/cobra"
)

func passiveCmd() *cobra.Command {
	passiveCMD := &cobra.Command{
		Use:     "passive",
		Short:   "Passive mode",
		Long:    `Passive mode`,
		Example: `regoftw passive`,
		Run:     runPassive,
	}

	return passiveCMD
}

func runPassive(cmd *cobra.Command, args []string) {
	// TEST
	urls := [][]string{
		{"url1",
			"test1", " hello"},
		{"url2",
			"test2", "hello tester"}}
	utils.ExecuteOnlineFunctions(urls)
}

func init() {
	rootCmd.AddCommand(passiveCmd())
}
