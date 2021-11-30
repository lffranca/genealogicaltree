package main

import (
	"context"
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

	ctx := context.Background()
	if err := clientGraphDB.Init.InsertData(ctx); err != nil {
		log.Panicln(err)
	}

	log.Println("importing successfully")
}

func init() {
	for _, envVar := range []string{
		"GRAPH_DB_HOST",
		"GRAPH_DB_PORT",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Panicf("Required enviroment variable not set: %s\n", envVar)
		}
	}
}
