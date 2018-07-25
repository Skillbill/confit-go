//
// command line utility for confit
//
package main

import (
	"github.com/Skillbill/confit-go"

	"flag"
	"fmt"
	"os"
)

var repoSecret string
var isAlias bool

func init() {
	flag.StringVar(&repoSecret, "s", "", "repo secret")
	flag.BoolVar(&isAlias, "a", false, "use alias")
	flag.Usage = func() {
		fmt.Printf("usage: %s [-s secret] [-a] repoId resource\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	repoId := flag.Arg(0)
	rsc := flag.Arg(1)

	if repoId == "" || rsc == "" {
		flag.Usage()
		os.Exit(1)
	}
	client := confit.Client{RepoId: repoId, Secret: repoSecret}
	load := client.LoadByPath
	if isAlias {
		load = client.LoadByAlias
	}
	p, err := load(rsc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load %s: %s\n", rsc, err)
		os.Exit(1)
	}
	fmt.Println(string(p))
}
