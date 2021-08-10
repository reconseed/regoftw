package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regoftw/conf"
	"regoftw/utils"

	"github.com/spf13/cobra"
)

func installCmd() *cobra.Command {
	installCMD := &cobra.Command{
		Use:     "install",
		Short:   "Install requirements",
		Long:    `Install requirements`,
		Example: `regoftw install`,
		Run:     runInstall,
	}
	return installCMD
}

func runInstall(cmd *cobra.Command, args []string) {
	installPath := conf.REGOPATH
	if utils.ExistFolder(installPath) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Installation folder found, do you want to overwrite the previous installation? (y/n): ")
		answer, _ := reader.ReadString('\n')
		if answer[0] != 89 && answer[0] != 121 {
			utils.PrintError("Aborting...")
			os.Exit(0)
		}
		utils.PrintInfo("Overwriting previous installation...")
	}
	utils.DeleteFolder(installPath)
	utils.PrintInfo("Installing data in " + installPath)
	utils.CreateDirectory(installPath)
	utils.PrintOK("Installation done!")
}

func init() {
	rootCmd.AddCommand(installCmd())
}
