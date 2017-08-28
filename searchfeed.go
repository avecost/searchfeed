package main

import (
	"log"
	"os"
	_ "github.com/avecost/searchfeed/matchers"
	"github.com/avecost/searchfeed/search"
)

func init() {
	// change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Harvey")
}
