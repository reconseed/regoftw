package cmd

import (
	"regoftw/core"

	"github.com/spf13/cobra"
)

var (
	urlRobots      string
	waybackmachine bool
	mode           uint
)

func roboxtractorCmd() *cobra.Command {
	roboxCMD := &cobra.Command{
		Use:     "roboxtractor [flags]",
		Short:   "Extract endpoints marked as disallow in robots.txt file",
		Long:    "Extract endpoints marked as disallow in robots.txt file.",
		Example: `regoftw roboxtractor -u https://example.com -wb`,
		Run:     runRoboxtractor,
	}

	roboxCMD.Flags().UintVarP(&mode, "mode", "m", 1, "Extract URLs (0) // Extract endpoints to generate a wordlist  (>1)")
	roboxCMD.Flags().StringVarP(&urlRobots, "url", "u", "", "URL to extract endpoints marked as disallow in robots.txt file")
	roboxCMD.Flags().BoolVar(&waybackmachine, "wb", false, "Check Wayback Machine. Check 5 years (Slow mode)")
	roboxCMD.MarkFlagRequired("domain")
	roboxCMD.Flags().SortFlags = false

	return roboxCMD
}

func runRoboxtractor(cmd *cobra.Command, args []string) {
	core.StartRoboxtractor(urlRobots, mode, waybackmachine)
}

func init() {
	rootCmd.AddCommand(roboxtractorCmd())
}
