package conf

import (
	"os"
)

type Context interface {
	GetVersion() string
	GetRegoPath() string
	GetWorkPlace() string
	GetDomains() []string
	IsIncremental() bool
	IsVerbose() bool
	SilentMode() bool
}

type context struct {
	version     string
	regoPath    string
	output      string
	domains     []string
	incremental bool
	verbose     bool
	silent      bool
}

var VERSION = "0.1b"
var REGOPATH = getHomePath() + "/regoftw"

func getHomePath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return dirname
}

var ctx Context = &context{
	version:     VERSION,
	regoPath:    REGOPATH,
	output:      "",
	domains:     nil,
	incremental: false,
	verbose:     false,
	silent:      false,
}

var generate = false

func GenerateCTX(output string, domains []string, incremental bool, verbose bool, silent bool) {
	if !generate {
		generate = true
		ctx = &context{
			version:     VERSION,
			regoPath:    REGOPATH,
			output:      output,
			domains:     domains,
			incremental: incremental,
			verbose:     verbose,
			silent:      silent,
		}
	}
}

func (ctx *context) GetVersion() string {
	return ctx.version
}

func (ctx *context) GetRegoPath() string {
	return ctx.regoPath
}

func (ctx *context) GetWorkPlace() string {
	return ctx.output
}

func (ctx *context) GetDomains() []string {
	return ctx.domains
}

func (ctx *context) IsIncremental() bool {
	return ctx.incremental
}

func (ctx *context) IsVerbose() bool {
	return ctx.verbose
}

func (ctx *context) SilentMode() bool {
	return ctx.silent
}

func GetCTX() Context {
	return ctx
}
