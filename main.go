package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var repoLink string

	flag.StringVar(&repoLink, "repoLink", "", "github ssh link to use")

	flag.Parse()
	// Check github
	if strings.Compare(repoLink, "") == 0 {
		fmt.Fprintf(os.Stdout, "enter a valid github repo link")
		os.Exit(0)
	}

}
