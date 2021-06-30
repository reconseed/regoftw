package core

import (
	"fmt"
	"regoftw/conf"
	"regoftw/utils"
	"strings"

	"github.com/barasher/go-exiftool"
)

func getMetadata(workPlace string, url string, res chan []exiftool.FileMetadata, guardThreads chan struct{}) {
	name := strings.Split(url, "/")
	docName := name[len(name)-1]
	fileName := workPlace + "/" + docName
	if !utils.DownloadFile(url, fileName) {
		utils.PrintErrorIfVerbose("Error downloading " + docName)
		res <- nil
		return
	}
	et, err := exiftool.NewExiftool()
	if err != nil {
		utils.PrintErrorIfVerbose("Error downloading " + docName)
		res <- nil
		return
	}
	defer et.Close()
	utils.PrintInfoIfVerbose(docName + " has been downloaded")
	fileInfo := et.ExtractMetadata(fileName)
	utils.DeleteFile(workPlace, docName)
	<-guardThreads
	res <- fileInfo
}

func ExtractMetadata(domain string, total int, threads int) {
	ctx := conf.GetCTX()
	workPlace := ctx.GetWorkPlace() + "/" + domain
	utils.CreateDirectory(workPlace)
	urls := utils.BingFiles(domain, total)
	var authors []string
	var software []string

	utils.PrintInfoIfVerbose(fmt.Sprintf("Total documents found: %d", len(urls)))
	guardThreads := make(chan struct{}, threads)
	for _, url := range urls {
		guardThreads <- struct{}{}
		result := make(chan []exiftool.FileMetadata, 1)
		go getMetadata(workPlace, url, result, guardThreads)
		fileInfos := <-result
		if fileInfos == nil {
			continue
		}
		for _, fileInfo := range fileInfos {
			if fileInfo.Err != nil {
				continue
			}
			for k, v := range fileInfo.Fields {
				if k == "Producer" { // Only software or authors
					new := fmt.Sprintf("%v", v)
					software = append(software, new)
				} else if k == "Author" {
					new := fmt.Sprintf("%v", v)
					authors = append(authors, new)
				}
			}
		}
	}

	utils.PrintInfoIfVerbose(fmt.Sprintf("Total authors found: %d", len(authors)))
	utils.PrintInfoIfVerbose(fmt.Sprintf("Total software found: %d", len(software)))
	utils.PrintInfoIfVerbose("Check path: " + workPlace)

	utils.WriteFileFromList(workPlace, "metadata_authors.txt", authors)
	utils.WriteFileFromList(workPlace, "metadata_software.txt", software)
	utils.PrintOK("Done!")
}
