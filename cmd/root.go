package cmd

import (
	"regoftw/conf"
	"regoftw/utils"

	"github.com/spf13/cobra"
)

var (
	workplace string
	verbose   bool
	silent    bool
	rootCmd   = &cobra.Command{
		Use:   "regoftw [options]",
		Short: "regoftw - Recon Tool",
		Long:  "regoftw - Awesome Recon Tool",
		Example: `regoftw recon -w /tmp/test --passive
regoftw emailfinder -w /tmp/test -d example.com
regoftw gotator -w /tmp/test -d domainx.txt -p permutations.txt
...`,
		Version: conf.GetCTX().GetVersion(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			UpdateConfig()
			utils.Banner()
		},
	}
)

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.PersistentFlags().StringVarP(&workplace, "workplace", "w", "", "reconGOFTW WorkPlace (mandatory)")
	rootCmd.MarkPersistentFlagRequired("workplace")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "regoftw doesn't show banner")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = false
}

func UpdateConfig() {
	conf.GenerateCTX(workplace, verbose, silent)
	if !utils.ExistFolder(workplace) && !utils.CreateDirectory(workplace) {
		panic("[-] Workplace is not found and it cannot be created.")
	}
}

func Execute() error {
	return rootCmd.Execute()
}
