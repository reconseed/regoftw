package utils

import (
	"bufio"
	"fmt"
	"os"
	"regoftw/conf"

	"github.com/tcnksm/go-latest"
)

func CheckLatestVersion() {
	PrintInfoIfVerbose("Checking the latest version in the repository")
	isLatest := true
	githubTag := &latest.GithubTag{
		Owner:      "reconseed",
		Repository: "regoftw",
	}

	res, err := latest.Check(githubTag, conf.VERSION)
	if err != nil {
		PrintError("Impossible to check latest version")
		return
	}
	if res.Outdated {
		msg := fmt.Sprintf("%s is not latest, you should upgrade to %s", conf.VERSION, res.Current)
		PrintError(msg)
		fmt.Println("|_ Run: go get -u github.com/reconseed/regoftw")
		isLatest = false
	}
	if !isLatest {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to continue the execution? (y/n): ")
		answer, _ := reader.ReadString('\n')
		if answer[0] != 89 && answer[0] != 121 {
			PrintError("Aborting...")
			os.Exit(0)
		}
	} else {
		PrintOKIfVerbose("regoFTW is in the latest version")
	}
}
