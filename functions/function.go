package functions

import (
	"regoftw/utils"
)

type DataFunction struct {
	app         string
	domain_flag string
	pipe        string
	flags       string
	output      []string
}

// data := DataFunction{app: "httpx", pipe: "cat file.txt", flags: "-status-code", output: []string{"f", "httpx.txt"}}
func RunFunction(domain string, data DataFunction) {
	if data.app == "" {
		//ERROR
	}
	dbmanager := utils.GetDBManager()
	if !dbmanager.CanRunFunction(domain, data.app) {
		return
	}
	dbmanager.UpdateStatus(domain, data.app, 1)
	utils.PrintInfoIfVerbose("Running " + data.app)

	command := ""
	if data.pipe != "" {
		command = data.pipe + " | " + data.app + data.flags
	} else {
		command = data.app + " " + data.domain_flag + " " + domain + " " + data.flags
	}
	if data.output[0] == "d" {
		// create directory
	}

	utils.ExecuteCommand(command)

	dbmanager.UpdateStatus(domain, data.app, 2)
	utils.PrintOKIfVerbose(data.app + " Done")
}

// domain := "target.com"
// params := "-d " + domain + " -all"
// RunFunction(domain, "subfinder", params, "[0, 'output.txt']")
// {"app": "subfinder", "domain_flag": "-d", "flags": "-all", "output": ["f", "lkkl"]}

// gauplus example.com -t 10 -random-agent -subs
