package utils

import (
	"io/ioutil"
	"os/exec"
)

//Execute one single command with pipes or not (Ex. `ls -la` OR `cat file.txt | wc -l`)
func ExecuteCommand(command string) (string, error) {
	out, err := exec.Command("bash", "-c", command).Output()
	output := ""
	if err == nil {
		output = string(out)
	}
	return output, err
}

// Read bash files and execute them with GO routines
func ExecuteBashFunctions(files []string) {
	var codes []string
	for _, file := range files {
		code, err := ioutil.ReadFile(file)
		if err == nil {
			codes = append(codes, string(code))
		}
	}
	ExecuteTasks(codes)
}
