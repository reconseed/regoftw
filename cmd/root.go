package cmd

import (
	"os"
	"regoftw/conf"
	"regoftw/functions"
	"regoftw/utils"
	"strings"

	"github.com/spf13/cobra"
)

// Check to add this vars in ctx or similar
var (
	workplace         string
	domain            string
	domainsFile       string
	excludeSubdomains string
	configFile        string // TODO: Generate the configuration file in install mode and use that path?
	domainsToCheck    []string
	incremental       bool
	verbose           bool
	silent            bool
	rootCmd           = &cobra.Command{
		Use:   "regoftw [options]",
		Short: "regoftw - Recon Tool",
		Long:  "regoftw - Awesome Recon Tool",
		Example: `regoftw active -w /tmp/test -d example.com
regoftw full -w /tmp/test -d example.com
regoftw recon -w /tmp/test -d example.com
regoftw passive -v -w /tmp/test -D ./targets.txt
...`,
		Version: conf.GetCTX().GetVersion(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			utils.Banner()
			utils.CheckLatestVersion()
			if cmd.Name() != "install" {
				if !utils.ExistFolder(conf.REGOPATH) {
					utils.PrintError(`Installation folder not found
|__ Please run: regoftw install`)
					os.Exit(0)
				}
				UpdateConfig()
			}
		},
	}
)

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.PersistentFlags().StringVarP(&workplace, "output", "o", conf.REGOPATH+"/reports", "regoFTW WorkPlace.")
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Domain to analyze")
	rootCmd.PersistentFlags().StringVarP(&domainsFile, "domains", "D", "", "File with domains to analyze. Absolute path or local path starting with ./")
	rootCmd.PersistentFlags().StringVarP(&excludeSubdomains, "exclude", "x", "", "File with domains to exclude from scope. Absolute path or local path starting with ./")
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "c", "", "Configuration file. Absolute path or local path starting with ./")
	rootCmd.PersistentFlags().BoolVarP(&incremental, "incremental", "i", false, "If a previous scanner exists, add any new data found.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "regoFTWdoesn't show banner")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = false
}

func generateDomains() {
	if domain == "" && domainsFile == "" {
		utils.PrintError(`You must specify either --domain (-d) or --domains (-D).
    An example -> regoftw passive -v -w /tmp/test -D ./targets.txt
		`)
		os.Exit(1)
	}
	if domain != "" {
		domainsToCheck = append(domainsToCheck, domain)
	} else {
		fileContent := utils.ReadFile(domainsFile)
		if fileContent == "" {
			os.Exit(1)
		}
		for _, line := range strings.Split(fileContent, "\n") {
			line = strings.Trim(line, " ")
			if line == "" {
				continue
			}
			// TODO: Improve this check for a valids (sub)domain
			if len(strings.Split(line, ".")) < 2 {
				utils.PrintInfoIfVerbose("Skipping bad (sub)domian " + line)
				continue
			}
			domainsToCheck = append(domainsToCheck, line)
		}
	}
	if len(domainsToCheck) == 0 {
		utils.PrintError("No valid (sub)domains found")
		os.Exit(1)
	}
}

func generateFolders() {
	conf.GenerateCTX(workplace, domainsToCheck, incremental, verbose, silent)
	if !utils.ExistFolder(workplace) && !utils.CreateDirectory(workplace) {
		utils.PrintError("Workplace is not found and it cannot be created.")
		os.Exit(1)
	}
	dbManager := utils.GetDBManager()
	for _, d := range domainsToCheck {

		outputDomain := workplace + "/" + d
		if !utils.ExistFolder(outputDomain) && !utils.CreateDirectory(outputDomain) {
			utils.PrintError("Domain Workplace (" + outputDomain + ") is not found and it cannot be created.")
			os.Exit(1)
		}
		if dbManager.ExistDomain(d) {
			utils.PrintInfoIfVerbose("Domain " + d + " exists in DB inside the output directory")
		} else {
			dbManager.GenerateDataDomain(d)
		}
	}
}

func generateConfig() {
	// Generate the configuration file in install mode and use that path?
	noConfigFile := true
	if configFile != "" {
		if utils.ExistFile(domainsFile) {
			conf.GenerateConfiguration(configFile)
			noConfigFile = false
		}

	}
	if noConfigFile {
		// TODO: Generate default configuration
	}
}

func UpdateConfig() {
	generateDomains()
	generateFolders()
	generateConfig()
	functions.GenerateResolvers()
}

func Execute() error {
	return rootCmd.Execute()
}
