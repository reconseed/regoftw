package cmd

import (
	"regoftw/core"

	"github.com/spf13/cobra"
)

var (
	emailTotal  int
	emailDomain string
)

func emailfinderCmd() *cobra.Command {
	emailCMD := &cobra.Command{
		Use:   "emailfinder [flags]",
		Short: "Search emails through Bing",
		Long: `Search emails through Bing. The objective is to extract 'real' emails 
		that are present in the search engine through the search '@domain.com'.`,
		Example: `regoftw emailfinder -d example.com -t 250`,
		Run:     runEmailFinder,
	}

	emailCMD.Flags().IntVarP(&emailTotal, "total", "t", 100, "Total results to check (20-350)")
	emailCMD.Flags().StringVarP(&emailDomain, "domain", "d", "", "Domain to search (mandatory)")
	emailCMD.MarkFlagRequired("domain")
	emailCMD.Flags().SortFlags = false

	return emailCMD
}

func runEmailFinder(cmd *cobra.Command, args []string) {
	if metaTotal < 20 {
		metaTotal = 20
	} else if metaTotal > 350 {
		metaTotal = 350
	}
	core.ExtractEmails(emailDomain, emailTotal)
}

func init() {
	rootCmd.AddCommand(emailfinderCmd())
}
