package utils

import (
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

func ExistFile(path string, fileName string) bool {
	exist := true
	file := joinPath(path, fileName)
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateFile(path string, fileName string) bool {
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

func DeleteFile(path string, fileName string) bool {
	success := true
	err := os.Remove(joinPath(path, fileName))
	if err != nil {
		success = false
	}
	return success
}

func ExistFolder(path string) bool {
	return ExistFile(path, "")
}

func CreateDirectory(path string) bool {
	success := true
	err := os.MkdirAll(path, 0766)
	if err != nil {
		success = false
	}
	return success
}

func WriteFile(path string, fileName string, data string) bool {
	var file *os.File
	var err error
	if !ExistFile(path, fileName) {
		if !CreateFile(path, fileName) {
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
	if !ExistFile(path, fileName) {
		if !CreateFile(path, fileName) {
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
