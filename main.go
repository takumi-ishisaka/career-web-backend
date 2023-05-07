package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nokin-all-of-career/career-web-backend/server"
)

const (
	// DEV : development
	DEV = 1
	// PRO : production
	PRO = 2
)

func main() {
	runOptionString := flag.String("runOption", "", "please select runOption, -runOption=DEV or -runOption=PRO")
	flag.Parse()

	runOption := DEV
	if *runOptionString == "PRO" {
		runOption = PRO
	} else if *runOptionString == "DEV" {
		runOption = DEV
	}
	err := server.Run(runOption)
	if err != nil {
		log.Fatal(fmt.Sprintf(`{"error":"%v"}`, err))
	}
}
