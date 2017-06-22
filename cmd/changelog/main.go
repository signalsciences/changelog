package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/signalsciences/changelog"
)

func main() {
	var (
		flagLastVersion  = flag.Bool("last-version", false, "Show last version only")
		flagLastEntry    = flag.Bool("last-entry", false, "Show last entry only")
		flagNoUnreleased = flag.Bool("no-unreleased", false, "Error if an unreleased section is present")

		flagNoComment = flag.Bool("no-comments", false, "Error if HTML comments are found")
	)

	flag.Parse()

	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	cl, err := changelog.Parse(string(raw))
	if err != nil {
		log.Fatal(err)
	}

	if *flagNoComment && bytes.Contains(raw, []byte("<!--")) {
		log.Fatalf("Found html comment")
	}
	if *flagNoUnreleased && cl.Unreleased != "" {
		log.Fatalf("Found unreleased section")
	}

	if *flagLastVersion {
		v := cl.Top().Version
		if v == "" {
			log.Fatalf("no version")
		}
		fmt.Println(v)
		return
	}
	if *flagLastEntry {
		e := cl.Top()
		fmt.Println(e)
		return
	}
	fmt.Println(cl.String())
}
