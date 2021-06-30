package conf

type Context interface {
	GetVersion() string
	GetWorkPlace() string
	IsVerbose() bool
	SilentMode() bool
}

type context struct {
	version   string
	workplace string
	verbose   bool
	silent    bool
}

var VERSION = "0.1b"

var ctx Context = &context{
	version:   VERSION,
	workplace: "",
	verbose:   false,
	silent:    false,
}

var generate = false

func GenerateCTX(workplace string, verbose bool, silent bool) {
	if !generate {
		generate = true
		ctx = &context{
			version:   VERSION,
			workplace: workplace,
			verbose:   verbose,
			silent:    silent,
		}
	}
}

func (ctx *context) GetVersion() string {
	return ctx.version
}

func (ctx *context) GetWorkPlace() string {
	return ctx.workplace
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
