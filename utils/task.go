package utils

import (
	"fmt"
	"sync"
)

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
