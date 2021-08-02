package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"sync"
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

// TODO: Check whether it is a better approach to run the functions like this or to create functions in GO
// Maybe GO functions in order to be able to better update the DB
// urls must provide a list of [URL, funcion name, args ]
func ExecuteOnlineFunctions(urls [][]string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		// Check if we need to execute the function
		// If the functions is executed... Update DB status 1
		wg.Add(1)
		script := GetRawResponse(url[0])
		go executeTask(script+"\n "+url[1]+" "+url[2], &wg)
		// If the function is executed... Update DB status 2 or -1
	}
	wg.Wait()

}

// data is [URL, funcion name, args ]
func ExecuteSingleOnlineFunction(data []string) {
	// Check if we need to execute the function
	// If the functions is executed... Update DB status 1
	script := GetRawResponse(data[0])
	out, err := ExecuteCommand(script + "\n " + data[1] + " " + data[2])
	if err != nil {
		PrintErrorIfVerbose(err.Error())
	} else {
		fmt.Println(out)
	}
	// If the function is executed... Update DB status 2 or -1
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

func executeTask(task string, wg *sync.WaitGroup) {
	defer wg.Done()
	output, err := ExecuteCommand(task)
	//TODO: Change result action
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println(output)
	}
}

func ExecuteTasks(tasks []string) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go executeTask(task, &wg)
	}
	wg.Wait()
}

// TODO: Decide what type of function to use
type StringFunction struct {
	Function func(params []string)
	Args     []string
}

type IntFunction struct {
	Function func(params []int)
	Args     []int
}

type ExecuteInt struct {
	Functions []IntFunction
}

type ExecuteString struct {
	Functions []StringFunction
}

func (execute *ExecuteInt) Run() {
	var wg sync.WaitGroup
	for _, f := range execute.Functions {
		wg.Add(1)
		go func(f1 func(params []int), params []int) {
			executeIntTask(f1, params)
			wg.Done()
		}(f.Function, f.Args)
	}
	wg.Wait()
}

func (execute *ExecuteString) Run() {
	var wg sync.WaitGroup
	for _, f := range execute.Functions {
		wg.Add(1)
		go func(f1 func(params []string), params []string) {
			executeStringTask(f1, params)
			wg.Done()
		}(f.Function, f.Args)
	}
	wg.Wait()
}

// Launch various functions with its arguments
func ExecuteGOListTask(execute ExecuteString) {
	var wg sync.WaitGroup
	for _, f := range execute.Functions {
		wg.Add(1)
		go func(f1 func(params []string), params []string) {
			executeStringTask(f1, params)
			wg.Done()
		}(f.Function, f.Args)
	}
	wg.Wait()
}

func executeStringTask(f func(params []string), params []string) {
	f(params)
}

func executeIntTask(f func(params []int), params []int) {
	f(params)
}
