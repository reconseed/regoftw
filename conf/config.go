package conf

import (
	"encoding/json"
	"io/ioutil"
)

type Config interface {
	GetReconConf() Configuration
}

type reconCFG struct {
	conf Configuration
}

var cfg Config

func GenerateConfiguration(file string) Config {
	if cfg == nil {
		cfg = newConfig(file)
	}
	return cfg
}

func newConfig(file string) Config {
	byteValue, _ := ioutil.ReadFile(file)
	var reconConf Conf
	json.Unmarshal(byteValue, &reconConf)
	return &reconCFG{reconConf.Configuration}
}

func (cfg *reconCFG) GetReconConf() Configuration {
	return cfg.conf
}

type Conf struct {
	Configuration Configuration `json:"configuration"`
}

type Configuration struct {
	General         General
	Golang_vars     Golangvars
	Tools           Tools
	Osint           Osint
	Subdomains      Subdomains
	Webdetection    Webdetection
	Webanalytics    Webanalytics
	Vulnerabilities Vulnerabilities
	Extrafeautres   Extrafeautres
	Httpoptions     Httpoptions
	Threads         Threads
	Timeouts        Timeouts
	Wordlists       Wordlists
	Axiomfleet      Axiomfleet
}

type General struct {
	TOOLS            string `json:"TOOLS"`
	TOOLPATH         string `json:"TOOLPATH"`
	UPDATE_RESOLVERS bool   `json:"UPDATE_RESOLVERS"`
	PROXY_URL        string `json:"PROXY_URL"`
	DIF              string `json:"DIF"`
}

type Golangvars struct {
	GOROOT string `json:"GOROOT"`
	GOPATH string `json:"GOPATH"`
	PATH   string `json:"PATH"`
}

type Tools struct {
	AMASS_CONFIG  string `json:"AMASS_CONFIG"`
	GITHUB_TOKENS string `json:"GITHUB_TOKENS"`
}

type Osint struct {
	OSINT            bool `json:"OSINT"`
	GOOGLE_DORKS     bool `json:"GOOGLE_DORKS"`
	GITHUB_DORKS     bool `json:"GITHUB_DORKS"`
	METADATA         bool `json:"METADATA"`
	EMAILS           bool `json:"EMAILS"`
	DOMAIN_INFO      bool `json:"DOMAIN_INFO"`
	METAFINDER_LIMIT int  `json:"METAFINDER_LIMIT"`
}

type Subdomains struct {
	SUBCRT                bool `json:"SUBCRT"`
	SUBANALYTICS          bool `json:"SUBANALYTICS"`
	SUBBRUTE              bool `json:"SUBBRUTE"`
	SUBSCRAPING           bool `json:"SUBSCRAPING"`
	SUBPERMUTE            bool `json:"SUBPERMUTE"`
	SUBTAKEOVER           bool `json:"SUBTAKEOVER"`
	SUBRECURSIVE          bool `json:"SUBRECURSIVE"`
	SUB_RECURSIVE_PASSIVE bool `json:"SUB_RECURSIVE_PASSIVE"`
	ZONETRANSFER          bool `json:"ZONETRANSFER"`
	S3BUCKETS             bool `json:"S3BUCKETS"`
}

type Webdetection struct {
	WEBPROBESIMPLE          bool   `json:"WEBPROBESIMPLE"`
	WEBPROBEFULL            bool   `json:"WEBPROBEFULL"`
	WEBSCREENSHOT           bool   `json:"WEBSCREENSHOT"`
	UNCOMMON_PORTS_WEB      string `json:"UNCOMMON_PORTS_WEB"`
	AXIOM_SCREENSHOT_MODULE string `json:"AXIOM_SCREENSHOT_MODULE"`
}

type Host struct {
	FAVICON          bool `json:"FAVICON"`
	PORTSCANNER      bool `json:"PORTSCANNER"`
	PORTSCAN_PASSIVE bool `json:"PORTSCAN_PASSIVE"`
	PORTSCAN_ACTIVE  bool `json:"PORTSCAN_ACTIVE"`
	CLOUD_IP         bool `json:"CLOUD_IP"`
}

type Webanalytics struct {
	WAF_DETECTION bool `json:"WAF_DETECTION"`
	NUCLEICHECK   bool `json:"NUCLEICHECK"`
	URL_CHECK     bool `json:"URL_CHECK"`
	URL_GF        bool `json:"URL_GF"`
	URL_EXT       bool `json:"URL_EXT"`
	JSCHECKS      bool `json:"JSCHECKS"`
	PARAMS        bool `json:"PARAMS"`
	FUZZ          bool `json:"FUZZ"`
	CMS_SCANNER   bool `json:"CMS_SCANNER"`
	WORDLIST      bool `json:"WORDLIST"`
}

type Vulnerabilities struct {
	XSS           bool `json:"XSS"`
	CORS          bool `json:"CORS"`
	TEST_SSL      bool `json:"TEST_SSL"`
	OPEN_REDIRECT bool `json:"OPEN_REDIRECT"`
	SSRF_CHECKS   bool `json:"SSRF_CHECKS"`
	CRLF_CHECKS   bool `json:"CRLF_CHECKS"`
	LFI           bool `json:"LFI"`
	SSTI          bool `json:"SSTI"`
	SQLI          bool `json:"SQLI"`
	BROKENLINKS   bool `json:"BROKENLINKS"`
	SPRAY         bool `json:"SPRAY"`
	BYPASSER4XX   bool `json:"BYPASSER4XX"`
	COMM_INJ      bool `json:"COMM_INJ"`
}

type Extrafeautres struct {
	NOTIFICATION      bool `json:"NOTIFICATION"`
	SOFT_NOTIFICATION bool `json:"SOFT_NOTIFICATION"`
	DEEP              bool `json:"DEEP"`
	DEEP_LIMIT        int  `json:"DEEP_LIMIT"`
	DIFF              bool `json:"DIFF"`
	REMOVETMP         bool `json:"REMOVETMP"`
	REMOVELOG         bool `json:"REMOVELOG"`
	PROXY             bool `json:"PROXY"`
	SENDZIPNOTIFY     bool `json:"SENDZIPNOTIFY"`
	PRESERVE          bool `json:"PRESERVE"`
}

type Httpoptions struct {
	HEADER string `json:"HEADER"`
}

type Threads struct {
	FFUF                   int `json:"FFUF"`
	HTTPX                  int `json:"HTTPX"`
	HTTPX_UNCOMMONPORTS    int `json:"HTTPX_UNCOMMONPORTS"`
	GOSPIDER               int `json:"GOSPIDER"`
	GITDORKER              int `json:"GITDORKER"`
	BRUTESPRAY_TH          int `json:"BRUTESPRAY_TH"`
	BRUTESPRAY_CONCURRENCE int `json:"BRUTESPRAY_CONCURRENCE"`
	ARJUN                  int `json:"ARJUN"`
	GAUPLUS                int `json:"GAUPLUS"`
	DALFOX                 int `json:"DALFOX"`
	PUREDNS_PUBLIC_LIMIT   int `json:"PUREDNS_PUBLIC_LIMIT"`
	PUREDNS_TRUSTED_LIMIT  int `json:"PUREDNS_TRUSTED_LIMIT"`
	DIRDAR                 int `json:"DIRDAR"`
	WEBSCREENSHOT          int `json:"WEBSCREENSHOT"`
	RESOLVE_DOMAINS        int `json:"RESOLVE_DOMAINS"`
}

type Timeouts struct {
	CMSSCAN             int `json:"CMSSCAN"`
	FFUF_MAXTIME        int `json:"FFUF_MAXTIME"`
	HTTPX               int `json:"HTTPX"`
	HTTPX_UNCOMMONPORTS int `json:"HTTPX_UNCOMMONPORTS"`
}

type Wordlists struct {
	FUZZ             string `json:"FUZZ"`
	LFI              string `json:"LFI"`
	SSTI             string `json:"SSTI"`
	SUBS             string `json:"SUBS"`
	SUBSBIG          string `json:"SUBSBIG"`
	RESOLVERS        string `json:"RESOLVERS"`
	TRUSTEDRESOLVERS string `json:"TRUSTEDRESOLVERS"`
}

type Axiomfleet struct {
	LAUNCH   bool   `json:"LAUNCH"`
	NAME     string `json:"NAME"`
	COUNT    int    `json:"COUNT"`
	REGIONS  string `json:"REGIONS"`
	SHUTDOWN bool   `json:"SHUTDOWN"`
}
