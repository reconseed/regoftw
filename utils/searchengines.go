package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

func hasSufix(href string) bool {
	extensions := []string{"pdf", "doc", "docx", "ppt", "pptx", "xls", "xlsx"}
	for _, ext := range extensions {
		if strings.HasSuffix(href, ext) {
			return true
		}
	}
	return false
}

func BingFiles(domain string, total int) []string {
	bigCount := 25
	if total < bigCount {
		bigCount = total
	}
	urlFirst := "https://www.bing.com/search?q=site:"
	urlSecond := "+(filetype:pdf+OR+filetype:doc+OR%20filetype:docx+OR+filetype:xls+OR+filetype:xlsx+OR+filetype:ppt+OR+filetype:pptx)&count="
	url := fmt.Sprintf("%s%s%s%d", urlFirst, domain, urlSecond, bigCount)
	var results []string
	iterCount := int(total / bigCount)
	if (total % bigCount) != 0 {
		iterCount += 1
	}
	var current int
	for count := 0; count < iterCount; count++ {
		current = count*bigCount + 1
		new_url := fmt.Sprintf("%s&first=%d&FORM=PERE", url, current)
		response := GetRawResponse(new_url)
		resultSoup := soup.HTMLParse(response)
		links := resultSoup.FindAll("a")
		for _, link := range links {
			href := link.Attrs()["href"]
			if strings.Contains(href, domain) && hasSufix(href) && ListContainsElement(results, href) == -1 {
				results = append(results, href)
				if len(results) >= total {
					break
				}
			}
		}
	}
	return results
}

func extractEmails(target string, data string) []string {
	pattern := regexp.MustCompile("[a-zA-Z0-9_\\.+-]+@" + target)
	data = strings.Replace(data, "<em>", "", -1)
	data = strings.Replace(data, "</em>", "", -1)
	data = strings.Replace(data, "<strong>", "", -1)
	data = strings.Replace(data, "</strong>", "", -1)
	data = strings.Replace(data, "<b>", "", -1)
	data = strings.Replace(data, "</b>", "", -1)
	aux := pattern.FindAllStringSubmatch(data, -1)
	var result []string
	for _, r := range aux {
		result = append(result, r[0])
	}
	return result
}

func BingEmails(domain string, total int) []string {
	var emails []string
	bigCount := 50
	if total < bigCount {
		bigCount = total
	}
	url := fmt.Sprintf("https://www.bing.com/search?q=inbody:'@%s'&count=%d", domain, bigCount)
	iterCount := int(total / bigCount)
	if (total % bigCount) != 0 {
		iterCount += 1
	}
	var current int
	for count := 0; count < iterCount; count++ {
		current = count*bigCount + 1
		new_url := fmt.Sprintf("%s&first=%d&FORM=PERE", url, current)
		response := GetRawResponse(new_url)
		results := extractEmails(domain, response)
		for _, email := range results {
			if ListContainsElement(emails, email) == -1 {
				emails = append(emails, email)
			}
		}
	}
	return emails
}
