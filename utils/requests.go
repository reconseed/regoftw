package utils

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func getURLResponse(url string) io.ReadCloser {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:89.0) Gecko/20100101 Firefox/89.0")
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	return res.Body
}

func GetRawResponse(url string) string {
	resBody := getURLResponse(url)
	if resBody != nil {
		defer resBody.Close()
		body, err := ioutil.ReadAll(resBody)
		if err != nil {
			return ""
		}
		return string(body)
	}
	return ""
}

func GetRequestReader(url string) io.ReadCloser {
	return getURLResponse(url) // Don't forget close it
}

func DownloadFile(url string, createFile string) bool {
	document := GetRequestReader(url)
	if document == nil {
		return false
	}
	defer document.Close()
	file, err := os.Create(createFile)
	if err != nil {
		return false
	}
	io.Copy(file, document)
	defer file.Close()
	return true
}
