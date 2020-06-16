package main

import (
	"bufio"
	"flag"
	"github.com/fatih/color"
	tld "github.com/jpillora/go-tld"
	"net/url"
	"os"
	"sync"
)

var unique_params []string
var unique_get []string
var unique_subs []string
var Threads int
var param bool
var path bool
var subs bool
var all bool
var output string
var path_result *os.File
var param_result *os.File
var sub_result *os.File

func ParseArguments() {
	flag.IntVar(&Threads, "t", 40, "Number of workers to use..default 40")
	flag.BoolVar(&param, "param", false, "If enabled..then get all params from stdin")
	flag.BoolVar(&path, "path", false, "If enabled..then get all paths from stdin")
	flag.BoolVar(&subs, "subs", false, "If enabled..then get all subs from stdin")
	flag.BoolVar(&all, "all", false, "If enabled..then get all paths,params,subs from stdin")
	flag.StringVar(&output, "o", " ", "Output to save as")
	flag.Parse()
}

func Banner() {
	color.HiCyan(`
        
      ____            .___.__.____    .__          __   
     /_   | ____    __| _/|__|    |   |__| _______/  |_ 
      |   |/    \  / __ | |  |    |   |  |/  ___/\   __\
      |   |   |  \/ /_/ | |  |    |___|  |\___ \  |  |  
      |___|___|  /\____ | |__|_______ \__/____  > |__|  
               \/      \/            \/       \/        1.0`)
	color.HiYellow("           1ndiList:- Recon Custom wordlist Generator")
	color.HiRed("              https://github.com/1ndianl33t")

	color.HiCyan("--------------------------------------------------------------")
}

func main() {
	Banner()
	ParseArguments()
	urls := make(chan string, Threads)
	processGroup := new(sync.WaitGroup)
	processGroup.Add(Threads)

	for i := 0; i < Threads; i++ {
		go func() {
			defer processGroup.Done()
			for u := range urls {
				get_things(u, param, path, subs, all, output)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		urls <- sc.Text()
	}
	close(urls)
	processGroup.Wait()
	print_unique()
}

func get_things(u string, param bool, path bool, subs bool, all bool, output string) {

	if all == true {
		get_path(u, output)
		get_param(u, output)
		get_subs(u, output)
	}

	if param == true {
		get_param(u, output)
	}

	if path == true {
		get_path(u, output)
	}
	if subs == true {
		get_subs(u, output)
	}

}

func get_param(u string, output string) {

	U, err := url.Parse(u)

	if err != nil {

	}

	m, _ := url.ParseQuery(U.RawQuery)

	for key, _ := range m {
		unique_params = append(unique_params, string(key))
	}
	if output != " " {
		if len(unique_params) != 0 {
			param_result, _ = os.Create(output + "_params.txt")
		}
	}

}

func get_path(u string, output string) {

	U, err := url.Parse(u)

	if err != nil {

	}

	g := U.EscapedPath()
	if g != "/" {
		unique_get = append(unique_get, g)
	}

	if output != " " {
		if len(unique_get) != 0 {
			path_result, _ = os.Create(output + "_paths.txt")
		}
	}
}

func get_subs(u string, output string) {

	U, err := tld.Parse(u)

	if err != nil {

	}

	subs := U.Subdomain
	if subs != "" {
		unique_subs = append(unique_subs, subs)
	}

	if output != " " {
		if len(unique_subs) != 0 {
			sub_result, _ = os.Create(output + "_subs.txt")
		}
	}

}

func unique(unique_vals []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range unique_vals {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func print_unique() {
	uni_para := unique(unique_params)
	if len(uni_para) != 0 {
		color.HiRed("[+] Total Unique Params Found %d", len(uni_para))
		for _, vals_para := range uni_para {
			color.HiCyan(vals_para)
			pr := bufio.NewWriter(param_result)
			pr.WriteString("\n" + vals_para)
			pr.Flush()
		}
	}

	uni_path := unique(unique_get)
	if len(uni_path) != 0 {
		color.HiRed("\n[+] Total Unique Paths Found %d", len(uni_path))
		for _, vals_path := range uni_path {
			color.HiGreen(vals_path)
			pt := bufio.NewWriter(path_result)
			pt.WriteString("\n" + vals_path)
			pt.Flush()
		}
	}
	uni_subs := unique(unique_subs)
	if len(uni_subs) != 0 {
		color.HiCyan("\n[+] Total Unique Subdomains Found %d", len(uni_subs))
		for _, vals_subs := range uni_subs {
			color.HiRed(vals_subs)
			sb := bufio.NewWriter(sub_result)
			sb.WriteString("\n" + vals_subs)
			sb.Flush()
		}
	}

}
