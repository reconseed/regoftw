package core

import (
	"fmt"
	"regoftw/conf"
	"regoftw/utils"
)

func ExtractEmails(domain string, total int) {
	ctx := conf.GetCTX()
	workPlace := ctx.GetWorkPlace() + "/" + domain

	utils.PrintInfoIfVerbose("Searching emails...")
	utils.CreateDirectory(workPlace)
	emails := utils.BingEmails(domain, total)
	utils.PrintInfoIfVerbose(fmt.Sprintf("Total Emails found: %d", len(emails)))
	utils.PrintInfoIfVerbose("Check path: " + workPlace)

	utils.WriteFileFromList(workPlace, "emails_bing.txt", emails)
	utils.PrintOK("Done!")
}
