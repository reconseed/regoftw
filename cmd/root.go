package cmd

import (
	"regoftw/conf"
	"regoftw/utils"

	"github.com/spf13/cobra"
)

var (
	workplace string
	domain    string
	verbose   bool
	silent    bool
	rootCmd   = &cobra.Command{
		Use:   "regoftw [options]",
		Short: "regoftw - Recon Tool",
		Long:  "regoftw - Awesome Recon Tool",
		Example: `regoftw active -w /tmp/test -d example.com
regoftw passive -w /tmp/test -d example.com
regoftw full -w /tmp/test -d example.com
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
	rootCmd.PersistentFlags().StringVarP(&workplace, "output", "o", "", "reconGOFTW WorkPlace (mandatory)")
	rootCmd.MarkPersistentFlagRequired("output")
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Domain to analyze")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "regoftw doesn't show banner")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = false
}

func UpdateConfig() {
	conf.GenerateCTX(workplace, domain, verbose, silent)
	if !utils.ExistFolder(workplace) && !utils.CreateDirectory(workplace) {
		panic("[-] Workplace is not found and it cannot be created.")
	}
}

func Execute() error {
	return rootCmd.Execute()
}
