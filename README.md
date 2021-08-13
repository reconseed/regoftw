<h1 align="center">
  <b>regoftw</b>
  <br>
</h1>
<p align="center">
  <a href="https://golang.org/dl/#stable">
    <img src="https://img.shields.io/badge/go-1.16-blue.svg?style=square&logo=go">
    
  </a>
   <a href="https://www.gnu.org/licenses/gpl-3.0.en.html">
    <img src="https://img.shields.io/badge/license-GNU-green.svg?style=square&logo=gnu">
  </a>
     <a href="https://github.com/reconseed/regoftw">
    <img src="https://img.shields.io/badge/version-0.01b-yellow.svg?style=square&logo=github">
  </a>
</p>
<p align="center">
 <a href="https://twitter.com/JosueEncinar">
    <img src="https://img.shields.io/badge/author-@JosueEncinar-orange.svg?style=square&logo=twitter">
  </a>
   <a href="https://twitter.com/six2dez1">
    <img src="https://img.shields.io/badge/author-@Six2dez1-orange.svg?style=square&logo=twitter">
  </a>
</p>


<p align="center">
regoftw is a recon tool ...
</p>
<br/>

# 🔨 Installation 

```
go get -u github.com/reconseed/regoftw
regoftw install
```

# ⚙️ Modes

With regoftw you can execute different modes:

* install: Install requirements. Necessary before using regoFTW for the first time.
* recon:  Full recon process (without attacks like sqli,ssrf,xss,ssti,lfi etc.)
* passive: Perform only passive steps
* full: Perform whole recon and all active attacks
* web: Perform only vulnerability checks/attacks on particular target
* subdomains: Perform only subdomain enumeration, web probing, subdomain takeovers
* osint: Performs an OSINT scan (no subdomain enumeration and attacks)


# 💫 Flags

**Global flags**

| Long flag | Short Flag | Description | Default Value |
|-----------|------------|-------------|---------------|
|  output | o  |   regoFTW output folder  |    ~/regoftw/reports      |
|  domain | d  |   Domain to analyze **\*** |    -      |
|  domains | D  |   File with domains to analyze. Absolute path or local path starting with ./ **\*** |  -     |
|  exclude | x |   File with domains to exclude from scope. Absolute path or local path starting with ./ |  -     |
|  conf | c |   Configuration file. Absolute path or local path starting with ./ **\*\***  |    -     |
| incremental | i | If a previous scanner exists, add any new data found. | false |
|  verbose | v    |    Verbose mode         |    false  |
|  silent | s   |  regoFTW doesn't show banner |    false  |
| version | | Show regoFTWversion | false |

*You need to define either **domain** or **domains**. If both are defined, only domain will be taken into account.
<br>\*\* If no configuration file is specified, a default one is used.

# 👾 Usage

```
regoftw <mode> -d target.com -v
regoftw <mode> -D domains.txt -i -v
```

# 🚀 Examples

```
regoftw full -d example.com -v -i
regoftw subdomains -d example.com
regoftw passive -D domains.txt -v
```

# 👨🏽‍💻 How to contribute

...

# 👉 Disclaimer

The tool is intended for use on targets with prior permission. The authors are not responsible for any illegitimate use.
