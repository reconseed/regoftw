package core

import (
	"fmt"
	"math/rand"
	"regoftw/conf"
	"regoftw/utils"
	"time"
)

func test1(test []int) { // TODO REMOVE
	value := rand.Intn(5)
	fmt.Println("BeginINT: ", test[value])
	time.Sleep(time.Duration(value) * time.Second)
	fmt.Println("EndINT: ", test[value])
}

func test2(test []string) { // TODO REMOVE
	value := rand.Intn(5)
	fmt.Println("BeginSTRING: " + test[value])
	time.Sleep(time.Duration(value) * time.Second)
	fmt.Println("EndSTRING: " + test[value])
}

// TODO: move to library
func StartRecon(active bool, passive bool, full bool, cfg string) {
	if active {
		fmt.Println("ACTIVE MODE")
	} else if full {
		fmt.Println("COMPLETE MODE")
	} else {
		fmt.Println("PASSIVE MODE")
	}
	workPlace := conf.GetCTX().GetWorkPlace()
	if cfg == "" {
		cfg = workPlace + "/ regofwt.json"
		utils.DownloadFile("URL", cfg)
		return //TODO: Add file config ur
	}
	configuration := conf.GenerateConfiguration(cfg).GetReconConf()

	fmt.Println(configuration.Osint)
	commands := []string{"pwd", "whoami", "ls | wc", "whoami"}
	utils.ExecuteTasks(commands)
	// osint.GetDomainInfo("as.com")

	fmt.Println(utils.ExistFile("/tmp/", "t2.txt"))
	utils.CreateFile("/tmp", "t2.txt")
	fmt.Println(utils.ExistFile("/tmp/", "t2.txt"))

	// args1 := []string{"1", "2", "3", "4", "5"}
	// var funcs utils.ExecuteString
	// function := utils.StringFunction{test2, args1}
	// funcs.Functions = append(funcs.Functions, function)
	// funcs.Functions = append(funcs.Functions, function)
	// funcs.Functions = append(funcs.Functions, function)
	// args1 = []string{"11", "22", "33", "44", "55"}
	// function = utils.StringFunction{test2, args1}
	// funcs.Functions = append(funcs.Functions, function)
	// funcs.Functions = append(funcs.Functions, function)
	// funcs.Run()

	// var funcs2 utils.ExecuteInt
	// args2 := []int{9, 8, 7, 6, 5}
	// function2 := utils.IntFunction{test1, args2}
	// funcs2.Functions = append(funcs2.Functions, function2)
	// funcs2.Functions = append(funcs2.Functions, function2)
	// funcs2.Functions = append(funcs2.Functions, function2)
	// funcs2.Run()

	// utils.ExecuteGOListTask(funcs)
	files := []string{"/tmp/data.sh", "/tmp/data2.sh"}
	utils.ExecuteBashFunctions(files)
	fmt.Println(utils.ExistFolder("/tmp/TEST2"))
	fmt.Println(utils.CreateDirectory("/tmp/TEST2"))
	fmt.Println(utils.CreateDirectory("/tmp/TEST2"))
	fmt.Println(utils.ExistFolder("/tmp/TEST2"))
	fmt.Println(workPlace)
	var mgr = utils.GetManager(workPlace)
	mgr.GenerateDataDomain("test.es")
	fmt.Println(mgr.CheckFunctionStatus("test.es", "gotator"))
	mgr.UpdateStatus("test.es", "gotator", 1)
	fmt.Println(mgr.CheckFunctionStatus("test.es", "gotator"))
}
