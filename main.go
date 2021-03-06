package main

import (
	"github.com/lffranca/genealogicaltree/pkg/ggin"
	"github.com/lffranca/genealogicaltree/pkg/ggogm"
	"log"
	"os"
	"strconv"
)

func main() {
	graphDBHost := os.Getenv("GRAPH_DB_HOST")
	graphDBPortS := os.Getenv("GRAPH_DB_PORT")

	graphDBPort, err := strconv.Atoi(graphDBPortS)
	if err != nil {
		log.Panicln(err)
	}

	clientGraphDB, err := ggogm.New(ggogm.Options{
		Host: &graphDBHost,
		Port: &graphDBPort,
	})
	if err != nil {
		log.Panicln(err)
	}

	defer func() {
		if err := clientGraphDB.Close(); err != nil {
			log.Println(err)
		}
	}()

	specPath := os.Getenv("SPEC_PATH")

	server, err := ggin.New(ggin.Options{
		SpecPath:         &specPath,
		PersonRepository: clientGraphDB.Person,
	})
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
		"GRAPH_DB_HOST",
		"GRAPH_DB_PORT",
		"ADDRESSES",
		"SPEC_PATH",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Panicf("Required enviroment variable not set: %s\n", envVar)
		}
	}
}
