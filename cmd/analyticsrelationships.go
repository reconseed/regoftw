package cmd

import (
	"regoftw/core"

	"github.com/spf13/cobra"
)

var (
	url string
)

func analyticsrelationshipsCmd() *cobra.Command {
	analyticsCMD := &cobra.Command{
		Use:   "analyticsrelationships [flags]",
		Short: "Search relationships in URLs",
		Long:  `Search relationships in URLs through Google Analytics ID.`,
		Example: `cat urls.txt | regoftw analyticsrelationships -v\n
		regoftw analyticsrelationships -v -u https://example.com`,
		Run: runAnalytics,
	}

	analyticsCMD.Flags().StringVarP(&url, "url", "u", "", "Url to check")
	analyticsCMD.Flags().SortFlags = false

	return analyticsCMD
}

func runAnalytics(cmd *cobra.Command, args []string) {
	if metaTotal < 20 {
		metaTotal = 20
	} else if metaTotal > 350 {
		metaTotal = 350
	}
	core.GetAnalyticsRelationships(url)
}

func init() {
	rootCmd.AddCommand(analyticsrelationshipsCmd())
}
