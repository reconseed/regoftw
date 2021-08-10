package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func joinPath(path string, fileName string) string {
	file := path
	if !strings.HasSuffix(file, "/") {
		file += "/"
	}
	file += fileName
	return file
}

func ExistFileInPath(path string, fileName string) bool {
	exist := true
	file := joinPath(path, fileName)
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ExistFile(fileName string) bool {
	exist := true
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateFileInPath(path string, fileName string) bool {
	success := true
	file := joinPath(path, fileName)
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		success = false
	} else {
		f.Close()
	}
	return success
}

func DeleteFileInPath(path string, fileName string) bool {
	success := true
	err := os.Remove(joinPath(path, fileName))
	if err != nil {
		success = false
	}
	return success
}

func DeleteFolder(path string) bool {
	success := true
	err := os.RemoveAll(path)
	if err != nil {
		success = false
	}
	return success
}

func ExistFolder(path string) bool {
	return ExistFileInPath(path, "")
}

func CreateDirectory(path string) bool {
	success := true
	err := os.MkdirAll(path, 0766)
	if err != nil {
		success = false
	}
	return success
}

func ReadFileInPath(path string, fileName string) string {
	if !ExistFileInPath(path, fileName) {
		PrintErrorIfVerbose("File " + fileName + " in " + path + " not found")
		return ""
	}
	data, err := ioutil.ReadFile(joinPath(path, fileName))
	if err != nil {
		return ""
	}
	return string(data)
}

func ReadFile(fileName string) string {
	if !ExistFile(fileName) {
		PrintErrorIfVerbose("File " + fileName + " not found")
		return ""
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return string(data)
}

func WriteFile(path string, fileName string, data string) bool {
	var file *os.File
	var err error
	if !ExistFileInPath(path, fileName) {
		if !CreateFileInPath(path, fileName) {
			return false
		}
		file, err = os.OpenFile(joinPath(path, fileName), os.O_RDWR, 0644)
	} else {
		file, err = os.OpenFile(joinPath(path, fileName), os.O_APPEND|os.O_WRONLY, 0644)
	}
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	if err != nil {
		return false
	}

	err = file.Sync()
	if err != nil {
		return false
	}
	return true
}

func WriteFileFromList(path string, fileName string, data []string) bool {
	var file *os.File
	var err error
	if !ExistFileInPath(path, fileName) {
		if !CreateFileInPath(path, fileName) {
			return false
		}
		file, err = os.OpenFile(joinPath(path, fileName), os.O_RDWR, 0644)
	} else {
		file, err = os.OpenFile(joinPath(path, fileName), os.O_APPEND|os.O_WRONLY, 0644)
	}
	if err != nil {
		return false
	}
	defer file.Close()
	for _, d := range data {
		_, err = file.WriteString(d + "\n")
		if err != nil {
			continue
		}
	}
	err = file.Sync()
	if err != nil {
		return false
	}
	return true
}
