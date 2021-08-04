package functions

import (
	"regoftw/conf"
	"regoftw/utils"
)

func PureDnsResolve(params []string) {
	//TODO: Create a list of filesnames
	//TODO: Check params
	//TODO: Create structures to manage data
	fileName := params[0]
	subsToResolve := params[1]
	resolvers := ""
	resolversTrusted := ""
	limit := ""
	rateLimit := ""
	logFile := ""
	output := conf.GetCTX().GetWorkPlace() + fileName
	utils.PrintInfoIfVerbose("Running PureDNS resolve")
	command := "puredns resolve " + subsToResolve + " -w " + output + " -r " + resolvers + " --resolvers-trusted " + resolversTrusted + " -l " + limit + " --rate-limit-trusted " + rateLimit + " 2>>" + logFile + " &>/dev/null"
	utils.ExecuteCommand(command)
	utils.PrintOKIfVerbose("PureDns reolve Done")
}

func PureDnsBrute(params []string) {
	utils.PrintInfoIfVerbose("Running PureDNS bruteforce")
	//command := "puredns bruteforce " +  + " " +  + " -w " +  + " -r " +  + " --resolvers-trusted " +  + " -l " +  + " --rate-limit-trusted " +  + " 2>>" +  + " &>/dev/null
	utils.PrintOKIfVerbose("PureDns bruteforce Done")
}
