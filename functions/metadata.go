package functions

import (
	"regoftw/utils"
)

func ExtractMetadata(params []string) {
	dbmanager := utils.GetDBManager()
	if !dbmanager.CanRunFunction(params[0], "metafinder") {
		return
	}
	dbmanager.UpdateStatus(params[0], "metafinder", 1)
	// Total and threads from configuration
	//	app.ExtractMetadata(params[0], 20, 20)
	dbmanager.UpdateStatus(params[0], "metafinder", 2)
}
