package cmd

import (
	"regoftw/core"

	"github.com/spf13/cobra"
)

var (
	flDomains            string
	flPermutations       string
	flDepth              uint
	flIterateNumbers     uint
	flPrefixes           bool
	flextractDomains     bool
	fladvancedOption     bool
	flminimizeDuplicates bool
	flThreads            uint
)

func gotatorCmd() *cobra.Command {
	gotatorCMD := &cobra.Command{
		Use:     "gotator [flags]",
		Short:   "DNS wordlists generator",
		Long:    `Gotator helps you to generate DNS wordlists through permutations.`,
		Example: `regoftw gotator -d domains.txt -p perms.txt -n 10`,
		Run:     runGotator,
	}

	gotatorCMD.Flags().StringVarP(&flDomains, "domains", "d", "", "List of domains to be swapped (1 per line) (mandatory)")
	gotatorCMD.MarkFlagRequired("domains")
	gotatorCMD.Flags().StringVarP(&flPermutations, "perm", "p", "", "List of permutations (1 per line)")
	gotatorCMD.Flags().UintVar(&flDepth, "depth", 1, "Specify the depth (Between 1 and 3")
	gotatorCMD.Flags().UintVarP(&flIterateNumbers, "numbers", "n", 0, "Permute the numbers found in the list of permutations")
	gotatorCMD.Flags().BoolVar(&flPrefixes, "prefixes", false, "Adding gotator prefixes to permutations")
	gotatorCMD.Flags().BoolVar(&flextractDomains, "md", false, "Extract domains and subdomains from subdomains found in 'sub' list")
	gotatorCMD.Flags().BoolVar(&fladvancedOption, "adv", false, "Advanced option. Generate permutations words with subdomains and words with -. And joins permutation word in the back (depth 1)")
	gotatorCMD.Flags().BoolVar(&flminimizeDuplicates, "mindup", false, "Set this flag to minimize duplicates. (For heavy workloads, it is recommended to activate this flag)")
	gotatorCMD.Flags().UintVarP(&flThreads, "threads", "t", 10, "Max Go routines")

	gotatorCMD.Flags().SortFlags = false

	return gotatorCMD
}

func runGotator(cmd *cobra.Command, args []string) {
	core.StartGotator(flDomains, flPermutations, flDepth, flIterateNumbers, flPrefixes, flextractDomains,
		fladvancedOption, flminimizeDuplicates, flThreads)
}

func init() {
	rootCmd.AddCommand(gotatorCmd())
}
