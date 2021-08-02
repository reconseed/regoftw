package conf

type Context interface {
	GetVersion() string
	GetWorkPlace() string
	GetDomain() string
	IsVerbose() bool
	SilentMode() bool
}

type context struct {
	version string
	output  string
	domain  string
	verbose bool
	silent  bool
}

var VERSION = "0.1b"

var ctx Context = &context{
	version: VERSION,
	output:  "",
	domain:  "",
	verbose: false,
	silent:  false,
}

var generate = false

func GenerateCTX(output string, domain string, verbose bool, silent bool) {
	if !generate {
		generate = true
		ctx = &context{
			version: VERSION,
			output:  output,
			domain:  domain,
			verbose: verbose,
			silent:  silent,
		}
	}
}

func (ctx *context) GetVersion() string {
	return ctx.version
}

func (ctx *context) GetWorkPlace() string {
	return ctx.output
}

func (ctx *context) GetDomain() string {
	return ctx.domain
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
