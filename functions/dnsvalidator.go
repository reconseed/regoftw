package functions

import (
	"regoftw/conf"
	"regoftw/utils"
)

func GenerateResolvers(params []string) {
	//TODO: Create a list of filesnames
	output := conf.GetCTX().GetWorkPlace() + "/resolvers.txt"
	utils.PrintInfoIfVerbose("Running GenerateResolvers")
	command := "dnsvalidator -tL https://public-dns.info/nameservers.txt -threads 100 -o " + output
	utils.ExecuteCommand(command)
	utils.PrintOKIfVerbose("GenerateResolvers Done")
}
