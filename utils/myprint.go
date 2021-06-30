package utils

import (
	"fmt"
	"regoftw/conf"
)

func PrintOK(msg string) {
	fmt.Println(GreenColor + "[+] " + ResetColor + msg)
}

func PrintInfo(msg string) {
	fmt.Println(YellowColor + "[i] " + ResetColor + msg)
}

func PrintError(msg string) {
	fmt.Println(RedColor + "[-] " + ResetColor + msg)
}

func PrintOKIfVerbose(msg string) {
	if conf.GetCTX().IsVerbose() {
		fmt.Println(GreenColor + "[+] " + ResetColor + msg)
	}
}

func PrintInfoIfVerbose(msg string) {
	if conf.GetCTX().IsVerbose() {
		fmt.Println(YellowColor + "[i] " + ResetColor + msg)
	}
}

func PrintErrorIfVerbose(msg string) {
	if conf.GetCTX().IsVerbose() {
		fmt.Println(RedColor + "[-] " + ResetColor + msg)
	}
}
