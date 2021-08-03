package cmd

import (
	"regoftw/functions"
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
	// // Online Option with bash functions
	// // TEST
	// urls := [][]string{
	// 	{"url1",
	// 		"test1", " hello"},
	// 	{"url2",
	// 		"test2", "hello tester"}}
	// utils.ExecuteOnlineFunctions(urls)
	// Go option
	var functionsToExecute utils.ExecuteRegoFunction
	function := utils.RegoFunction{Function: functions.GenerateResolvers, Args: nil}
	functionsToExecute.Functions = append(functionsToExecute.Functions, function)
	functionsToExecute.Run()
}

func init() {
	rootCmd.AddCommand(passiveCmd())
}
