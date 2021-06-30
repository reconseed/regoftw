package cmd

import (
	"regoftw/core"

	"github.com/spf13/cobra"
)

var (
	passive bool
	active  bool
	full    bool
	cfg     string
)

func reconCmd() *cobra.Command {
	reconCMD := &cobra.Command{
		Use:   "recon [flags]",
		Short: "Recon mode",
		Long:  `This module is the most complete, and allows to automate the entire recon process.`,
		Run:   runRecon,
	}

	reconCMD.Flags().BoolVarP(&passive, "passive", "p", true, "regoftw passive mode")
	reconCMD.Flags().BoolVarP(&active, "active", "a", false, "regoftw active mode")
	reconCMD.Flags().BoolVarP(&full, "full", "f", false, "regoftw complete mode")
	reconCMD.Flags().StringVarP(&cfg, "cfg", "c", "false", "Recon file configuration")
	reconCMD.Flags().SortFlags = false

	return reconCMD
}

func runRecon(cmd *cobra.Command, args []string) {
	core.StartRecon(active, passive, full, cfg)
}

func init() {
	rootCmd.AddCommand(reconCmd())
}
