package main

import (
	"github.com/lffranca/genealogicaltree/pkg/ggin"
	"log"
	"os"
)

func main() {
	specPath := os.Getenv("SPEC_PATH")

	server, err := ggin.New(ggin.Options{SpecPath: &specPath})
	if err != nil {
		log.Panicln(err)
	}

	addresses := os.Getenv("ADDRESSES")
	if err := server.Run(addresses); err != nil {
		log.Panicln(err)
	}
}

func init() {
	for _, envVar := range []string{
		"ADDRESSES",
		"SPEC_PATH",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Panicf("Required enviroment variable not set: %s\n", envVar)
		}
	}
}
