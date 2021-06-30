package cmd

import (
	"regoftw/core"

	"github.com/spf13/cobra"
)

var (
	metaTotal   int
	metaDomain  string
	metaThreads int
)

func metafinderCmd() *cobra.Command {
	metaCMD := &cobra.Command{
		Use:     "metafinder [flags]",
		Short:   "Search metadata through Bing",
		Long:    "Search metadata through Bing. The goal is to extract authors and software found in files like pdf or doc.",
		Example: `regoftw metafinder -d example.com -l 200`,
		Run:     runMetafinder,
	}

	metaCMD.Flags().IntVarP(&metaTotal, "limit", "l", 100, "Total results to check (5-300)")
	metaCMD.Flags().StringVarP(&metaDomain, "domain", "d", "", "Domain to search (mandatory)")
	metaCMD.Flags().IntVarP(&metaThreads, "threads", "t", 10, "Total Threads to download documents")
	metaCMD.MarkFlagRequired("domain")
	metaCMD.Flags().SortFlags = false

	return metaCMD
}

func runMetafinder(cmd *cobra.Command, args []string) {
	if metaTotal < 5 {
		metaTotal = 5
	} else if metaTotal > 300 {
		metaTotal = 300
	}
	core.ExtractMetadata(metaDomain, metaTotal, metaThreads)
}

func init() {
	rootCmd.AddCommand(metafinderCmd())
}
