
# 1ndiList v 1.0


*Take a list of URLs and returns their Paths , params , subs.*

***Warning: This is a work in progress.***

## Install

With Go:

```bash
▶ go get -u github.com/1ndianl33t/1ndiList

```
## Options:

```
▶ 1ndiList -h
Options:
  -subs, If enabled..then get all subs from stdin
  
  -param, If enabled..then get all params from stdin
 
  -path, If enabled..then get all paths from stdin

  -all, If enabled..then get all paths,params,subs from stdin
 
  -t, Number of workers to use..default 40 (default 40)
  
  -o, <dir> name   Directory to save  (name_params.txt)
```
## usage
```bash
▶ cat urls.txt | 1ndiList  -path -t 50 -o site

▶ cat urls.txt | 1ndiList  -param -t 50 -o site

▶ cat subdomainlist.txt | 1ndiList  -subs -t 50 -o site

▶ cat .txt | 1ndiList -all -t 50 -o site 
```
## Example
```bash

▶ assetfinder --subs-only test.com | waybackurls| 1ndiList -t 50 -all -o test.com

▶ Subfinder -d test.com -silent | httprobe | 1ndiList -subs -t 50 -o site.com

```
### Tips
```

```
## Output
```bash
      ____            .___.__.____    .__          __
     /_   | ____    __| _/|__|    |   |__| _______/  |_
      |   |/    \  / __ | |  |    |   |  |/  ___/\   __\
      |   |   |  \/ /_/ | |  |    |___|  |\___ \  |  |
      |___|___|  /\____ | |__|_______ \__/____  > |__|
               \/      \/            \/       \/        1.0
           1ndiList:- Recon Costume Wordlist Generator
              https://github.com/1ndianl33t
--------------------------------------------------------------

[+] Total Unique Paths Found 1728
/sitemap/keywords-pants-13-1-sitemap.xml
/sitemap/keywords-pretty-beauty-86-1-sitemap.xml
/sitemap/keywords-print-dresses-43-1-sitemap.xml
/sitemap/keywords-rings-27-1-sitemap.xml
/sitemap/keywords-sandals-26-1-sitemap.xml
/sitemap/keywords-scarves-35-1-sitemap.xml
/sitemap/keywords-shoes-sitemap.xml
/sitemap/keywords-shorts-39-1-sitemap.xml
```

### Donations
You can encourage me to contribute more to the open source with donations.

- Paypal - [https://www.paypal.me/1ndianl33t](https://www.paypal.me/1ndianl33t)

- GooglePay,Paytm -

`8085778875`

### Contributers

[![Twitter](https://img.shields.io/badge/twitter-@shivangx01b-blue.svg)](https://twitter.com/shivangx01b)

# Contact
[![Twitter](https://img.shields.io/badge/twitter-@1ndianl33t-blue.svg)](https://twitter.com/1ndianl33t)



