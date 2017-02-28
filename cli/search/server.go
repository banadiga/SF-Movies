package main

import (
	. "../../search-api"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var indexUrl string

func run(c *cli.Context) {
	log.Println("Creating search-api server...")
	config := NewConfig(indexUrl)

	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting search-api server...")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func migrate(c *cli.Context) {
	log.Println("Creating search-api server...")
	config := NewConfig(indexUrl)

	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting search-api migrate...")
	if err := server.Migrate(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "SF Movies service"
	app.Usage = "work with the `SF Movies` search api"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "index",
			Value: "localhost:9200",
			Usage: "path to index",
			EnvVar: "APP_SEARCH_INDEX",
			Destination: &indexUrl,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run the `SF Movies` search api",
			Action: run,
		},
		{
			Name:  "migrate",
			Usage: "Perform migrations",
			Action: migrate,
		},
	}

	app.Run(os.Args)
}
