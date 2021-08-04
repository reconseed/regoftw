package app

import (
	"fmt"
	"regexp"
	"regoftw/conf"
	"regoftw/utils"
	"strings"
	"time"

	"github.com/rjeczalik/wayback"
)

const (
	layoutISO = "2006-01-02"
)

var (
	roboxFile utils.File
)

func getDisallows(data string) [][]string {
	pattern := regexp.MustCompile("Disallow:\\s?.+")
	return pattern.FindAllStringSubmatch(data, -1)
}

func treatEndpoint(urlCheck string, entry string, endpoints []string, mode uint) []string {
	aux := strings.Split(entry, "Disallow:")
	if len(aux) <= 1 {
		return endpoints
	}
	endpoint := strings.Trim(aux[1], " ")
	if endpoint == "/" || endpoint == "*" || endpoint == "" {
		return endpoints
	}
	finalEndpoint := strings.Replace(endpoint, "*", "", -1)

	var finalPrint string
	for strings.HasPrefix(finalEndpoint, "/") {
		if len(finalEndpoint) >= 1 {
			finalEndpoint = finalEndpoint[1:] // Ex. /*/test or /*/*/demo
		} else {
			return endpoints
		}
	}
	for strings.HasSuffix(finalEndpoint, "/") {
		if len(finalEndpoint) >= 1 {
			finalEndpoint = finalEndpoint[0 : len(finalEndpoint)-1]
		} else {
			return endpoints
		}
	}
	if mode == 0 {
		finalPrint = urlCheck + "/" + finalEndpoint
	} else {
		finalPrint = finalEndpoint
	}

	if len(finalPrint) > 0 {
		if utils.ListContainsElement(endpoints, finalPrint) > -1 { // Avoid duplicates. Ex. view/ view/*
			return endpoints
		}
		endpoints = append(endpoints, finalPrint)
		roboxFile.WriteLn(finalPrint)
	}
	return endpoints
}

func waybackMachine(urlCheck string, endpoints []string, mode uint) {
	currentYear := time.Now().Year()
	robots := "/robots.txt"
	startYear := currentYear - 5 // Check last 5 years (It ignores current year)
	url := "https://web.archive.org/web/"
	lastURL := ""
	for startYear < currentYear {
		timestamp, err := wayback.ParseTimestamp(layoutISO, fmt.Sprintf("%04d-01-01", startYear))
		// wbMsg := fmt.Sprintf("%s. Wayback Machine Year %d", urlCheck, startYear)
		startYear += 1
		if err != nil {
			// utils.PrintInfoIfVerbose(fmt.Sprintf("WB %d - %s", startYear, err.Error()))
			continue
		}
		_, t, err := wayback.AvailableAt(urlCheck, timestamp)
		if err != nil {
			// utils.PrintInfoIfVerbose(fmt.Sprintf("WB %d - %s", startYear, err.Error()))
			continue
		}
		date := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		finalurl := url + date + "if_/" + urlCheck
		if finalurl == lastURL {
			// utils.PrintInfoIfVerbose(fmt.Sprintf("Skiping year %d. Same snapshot as previous", startYear-1))
			continue
		}
		lastURL = finalurl
		response := utils.GetRawResponse(finalurl + robots)

		endpoints = parseResponse(urlCheck, response, endpoints, mode)

	}
}

func thisWork(urlCheck string, mode uint, wayback bool) bool {
	var endpoints []string
	success := true
	robots := "/robots.txt"
	if strings.HasSuffix(urlCheck, "/") {
		urlCheck = urlCheck[0 : len(urlCheck)-1]
	}
	response := utils.GetRawResponse(urlCheck + robots)
	endpoints = parseResponse(urlCheck, response, endpoints, mode)
	if wayback {
		waybackMachine(urlCheck, endpoints, mode)
	}
	return success
}

func parseResponse(urlCheck string, response string, endpoints []string, mode uint) []string {
	allDisallows := getDisallows(response)
	if len(allDisallows) == 0 {
		// utils.PrintInfoIfVerbose("Nothing found here...")
		return endpoints
	}
	// utils.PrintInfoIfVerbose(fmt.Sprintf("Total entries marked as disallow: %d. Parsing and cleaning...", len(allDisallows)))
	for _, entry := range allDisallows {
		endpoints = treatEndpoint(urlCheck, entry[0], endpoints, mode)
	}
	return endpoints
}

func startJob(urlCheck string, mode uint, wayback bool) {
	if len(strings.Split(urlCheck, ".")) <= 1 {
		// utils.PrintErrorIfVerbose("URL format error " + urlCheck)
		return
	}
	if !strings.HasPrefix(urlCheck, "http") {
		if !(thisWork("https://"+urlCheck, mode, wayback)) {
			thisWork("http://"+urlCheck, mode, wayback)
		}
	} else {
		thisWork(urlCheck, mode, wayback)
	}
}

func StartRoboxtractor(url string, mode uint, waybackmachine bool) {
	ctx := conf.GetCTX()
	workPlace := ctx.GetWorkPlace()
	fileName := workPlace + "/roboxtractor-" + time.Now().Format("2006.01.01_15:04") + ".txt"
	roboxFile = utils.GetFile(fileName)
	startJob(url, mode, waybackmachine)
	roboxFile.Close()
	utils.PrintInfoIfVerbose("Check file: " + fileName + "\n")
	utils.PrintOKIfVerbose("Roboxtractor Done!")
}
