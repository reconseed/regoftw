package cmd

import (
	"os"
	"regoftw/conf"
	"regoftw/utils"
	"strings"

	"github.com/spf13/cobra"
)

// Check to add this vars in ctx or similar
var (
	workplace      string
	domain         string
	domainsFile    string
	configFile     string
	domainsToCheck []string
	verbose        bool
	silent         bool
	rootCmd        = &cobra.Command{
		Use:   "regoftw [options]",
		Short: "regoftw - Recon Tool",
		Long:  "regoftw - Awesome Recon Tool",
		Example: `regoftw active -w /tmp/test -d example.com
regoftw full -w /tmp/test -d example.com
regoftw active -w /tmp/test -d example.com
regoftw passive -v -w /tmp/test -D ./targets.txt
...`,
		Version: conf.GetCTX().GetVersion(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			utils.Banner()
			UpdateConfig()
		},
	}
)

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.PersistentFlags().StringVarP(&workplace, "output", "o", "", "reconGOFTW WorkPlace (mandatory)")
	rootCmd.MarkPersistentFlagRequired("output")
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Domain to analyze")
	rootCmd.PersistentFlags().StringVarP(&domainsFile, "domains", "D", "", "File with domains to analyze. Absolute path or local path starting with ./")
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "c", "", "Configuration file. Absolute path or local path starting with ./")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "regoftw doesn't show banner")
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
		data := strings.Split(domainsFile, "/")
		fileName := data[len(data)-1]
		path := strings.Join(data[0:len(data)-1], "/")
		fileContent := utils.ReadFile(path, fileName)
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
	conf.GenerateCTX(workplace, domain, verbose, silent)
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
			utils.PrintInfo("Domain " + d + " exists in DB inside the output directory")
		} else {
			dbManager.GenerateDataDomain(d)
		}
	}
}

func generateConfig() {
	noConfigFile := true
	if configFile != "" {
		splitConfig := strings.Split(domainsFile, "/")
		configName := splitConfig[len(splitConfig)-1]
		path := strings.Join(splitConfig[0:len(splitConfig)-1], "/")
		if utils.ExistFile(path, configName) {
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
}

func Execute() error {
	return rootCmd.Execute()
}
