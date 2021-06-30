package core

import (
	"bufio"
	"os"
	"regexp"
	"regoftw/conf"
	"regoftw/utils"
	"strings"
)

func getGoogleTagManager(targetURL string) (bool, []string) {
	var resultTagManager []string
	response := utils.GetRawResponse(targetURL)
	if response != "" {
		pattern := regexp.MustCompile("www\\.googletagmanager\\.com/ns\\.html\\?id=[A-Z0-9\\-]+")
		data := pattern.FindStringSubmatch(response)
		if len(data) > 0 {
			resultTagManager = append(resultTagManager, "https://"+strings.Replace(data[0], "ns.html", "gtm.js", -1))
		} else {
			pattern = regexp.MustCompile("GTM-[A-Z0-9]+")
			data = pattern.FindStringSubmatch(response)
			if len(data) > 0 {
				resultTagManager = append(resultTagManager, "https://www.googletagmanager.com/gtm.js?id="+data[0])
			} else {
				pattern = regexp.MustCompile("UA-\\d+-\\d+")
				aux := pattern.FindAllStringSubmatch(response, -1)
				var result []string
				for _, r := range aux {
					result = append(result, r[0])
				}
				return true, result
			}
		}
	}
	return false, resultTagManager
}

func getUA(url string) []string {
	pattern := regexp.MustCompile("UA-[0-9]+-[0-9]+")
	response := utils.GetRawResponse(url)
	var result []string
	if response != "" {
		aux := pattern.FindAllStringSubmatch(response, -1)
		for _, r := range aux {
			result = append(result, r[0])
		}
	} else {
		result = nil
	}
	return result
}

func cleanRelationShips(domains [][]string) []string {
	var allDomains []string
	for _, domain := range domains {
		allDomains = append(allDomains, strings.Replace(domain[0], "/relationships/", "", -1))
	}
	return allDomains
}

func getDomainsFromBuiltWith(id string) []string {
	pattern := regexp.MustCompile("/relationships/[a-z0-9\\-\\_\\.]+\\.[a-z]+")
	url := "https://builtwith.com/relationships/tag/" + id
	response := utils.GetRawResponse(url)
	var allDomains []string = nil
	if response != "" {
		allDomains = cleanRelationShips(pattern.FindAllStringSubmatch(response, -1))
	}
	return allDomains
}

func getDomainsFromHackerTarget(id string) []string {
	url := "https://api.hackertarget.com/analyticslookup/?q=" + id
	response := utils.GetRawResponse(url)
	var allDomains []string = nil
	if response != "" && !strings.Contains(response, "API count exceeded") {
		allDomains = strings.Split(response, "\n")
	}
	return allDomains
}

func getDomains(id string) []string {
	var allDomains []string = getDomainsFromBuiltWith(id)
	domains2 := getDomainsFromHackerTarget(id)
	if domains2 != nil {
		for _, domain := range domains2 {
			if utils.ListContainsElement(allDomains, domain) == -1 {
				allDomains = append(allDomains, domain)
			}
		}
	}
	return allDomains
}

func writeDomains(workPlace string, fileName string, ua string) {
	allDomains := getDomains(ua)
	if len(allDomains) == 0 {
		utils.WriteFile(workPlace, fileName, ua+" nothing found\n")
	} else {
		allDomains = append([]string{ua, "------------"}, allDomains...)
		utils.WriteFileFromList(workPlace, fileName, allDomains)
	}
}

func start(url string) {
	ctx := conf.GetCTX()
	workPlace := ctx.GetWorkPlace()
	verbose := ctx.IsVerbose()
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	if verbose {
		utils.PrintInfo("Analyzing url: " + url)
	}

	uaResult, resultTagManager := getGoogleTagManager(url)
	if len(resultTagManager) > 0 {
		var visited = []string{}
		var allUAs []string
		if !uaResult {
			urlGoogleTagManager := resultTagManager[0]
			if verbose {
				utils.PrintInfo("URL with UA: " + urlGoogleTagManager)
			}
			allUAs = getUA(urlGoogleTagManager)
		} else {
			if verbose {
				utils.PrintInfo("Found UA directly")
			}
			allUAs = resultTagManager
		}
		if verbose {
			utils.PrintInfo("Obtaining information from builtwith and hackertarget\n")
		}
		finalWorkPlace := strings.Replace(url, "http://", "", -1)
		finalWorkPlace = workPlace + "/" + strings.Replace(finalWorkPlace, "https://", "", -1)
		utils.CreateDirectory(finalWorkPlace)
		for _, ua := range allUAs {
			baseUA := strings.Join(strings.Split(ua, "-")[0:2], "-")
			if utils.ListContainsElement(visited, baseUA) == -1 {
				visited = append(visited, baseUA)
				writeDomains(finalWorkPlace, "analyticsRelationships.txt", baseUA)
			}
		}
		if verbose {
			utils.PrintInfo("Check Path: " + finalWorkPlace + "/analyticsRelationships.txt")
		}
		utils.PrintOK("Done!")
	} else {
		if verbose {
			utils.PrintInfo("Tagmanager URL not found\n")
		}
	}
}

func GetAnalyticsRelationships(url string) {
	if url != "" {
		start(url)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			start(scanner.Text())
		}
	}

}
