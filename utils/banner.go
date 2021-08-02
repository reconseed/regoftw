package utils

import (
	"math/rand"
	"regoftw/conf"
	"time"
)

func Banner() {
	if !conf.GetCTX().SilentMode() {
		data := GreenColor
		rand.Seed(time.Now().UTC().UnixNano())
		if rand.Intn(2) != 1 {
			data += `
	██████╗ ███████╗ ██████╗  ██████╗ ███████╗████████╗██╗    ██╗
	██╔══██╗██╔════╝██╔════╝ ██╔═══██╗██╔════╝╚══██╔══╝██║    ██║
	██████╔╝█████╗  ██║  ███╗██║   ██║█████╗     ██║   ██║ █╗ ██║
	██╔══██╗██╔══╝  ██║   ██║██║   ██║██╔══╝     ██║   ██║███╗██║
	██║  ██║███████╗╚██████╔╝╚██████╔╝██║        ██║   ╚███╔███╔╝
	╚═╝  ╚═╝╚══════╝ ╚═════╝  ╚═════╝ ╚═╝        ╚═╝    ╚══╝╚══╝ 

	`
			data += YellowColor + "\n"
			data += "  >> by @JosueEncinar && @six2dez1 \n"
		} else {
			data += `
	██▀███  ▓█████   ▄████  ▒█████    █████▒▄▄▄█████▓ █     █░
	▓██ ▒ ██▒▓█   ▀  ██▒ ▀█▒▒██▒  ██▒▓██   ▒ ▓  ██▒ ▓▒▓█░ █ ░█░
	▓██ ░▄█ ▒▒███   ▒██░▄▄▄░▒██░  ██▒▒████ ░ ▒ ▓██░ ▒░▒█░ █ ░█ 
	▒██▀▀█▄  ▒▓█  ▄ ░▓█  ██▓▒██   ██░░▓█▒  ░ ░ ▓██▓ ░ ░█░ █ ░█ 
	░██▓ ▒██▒░▒████▒░▒▓███▀▒░ ████▓▒░░▒█░      ▒██▒ ░ ░░██▒██▓ 
	░ ▒▓ ░▒▓░░░ ▒░ ░ ░▒   ▒ ░ ▒░▒░▒░  ▒ ░      ▒ ░░   ░ ▓░▒ ▒  
	  ░▒ ░ ▒░ ░ ░  ░  ░   ░   ░ ▒ ▒░  ░          ░      ▒ ░ ░  
	  ░░   ░    ░   ░ ░   ░ ░ ░ ░ ▒   ░ ░      ░        ░   ░  
	   ░        ░  ░      ░     ░ ░                       ░    
																													 
	`
			data += YellowColor + "\033[33m\n"
			data += "  >> by @six2dez1 && @JosueEncinar\n"
		}

		data += "  >> Version " + conf.GetCTX().GetVersion() + ResetColor + "\n"

		println(data)
	}
}
