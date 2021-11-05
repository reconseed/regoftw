package functions

import (
	"os"
	"regoftw/conf"
	"regoftw/utils"
	"time"
)

func GenerateResolvers() {
	//TODO: Create a list of filenames
	output := conf.GetCTX().GetWorkPlace() + "/resolvers.txt"
	// config := conf.Config.GetReconConf()
	file, err := os.Stat(output)
	if err != nil {
		utils.PrintInfoIfVerbose("We can't check resolvers date")
	} else {
		modifiedTime := file.ModTime()
		now := time.Now()
		if !now.After(modifiedTime.Add(24 * time.Hour)) {
			utils.PrintInfoIfVerbose("We found valid resolvers")
			return
		}
	}

	utils.PrintInfoIfVerbose("Generating resolvers")
	command := "dnsvalidator -tL https://public-dns.info/nameservers.txt -threads 100 -o " + output
	utils.ExecuteCommand(command)
	utils.PrintOKIfVerbose("GenerateResolvers Done")
}
