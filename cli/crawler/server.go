package main

import (
	. "../../crawler-api"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var autoUrl bool
var sfgovAppToken string
var indexType string
var indexUrl string

func run(context *cli.Context) {
	log.Println("Creating crawler server...")
	config := NewConfig(autoUrl, sfgovAppToken, indexType, indexUrl)

	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting crawler server...")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func migrate(context *cli.Context) {
	log.Println("Creating crawler server...")
	config := NewConfig(autoUrl, sfgovAppToken, indexType, indexUrl)

	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting crawler migrate...")
	if err := server.Migrate(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "SF Movies crawler"
	app.Usage = "work with the `SF Movies` crawler"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "auto-start",
			Usage: "auto start",
			EnvVar: "APP_CRAWLER_AUTO_START",
			Destination: &autoUrl,
		},
		cli.StringFlag{
			Name: "index-type",
			Value: "elasticsearch",
			Usage: "index type",
			EnvVar: "APP_CRAWLER_TYPE",
			Destination: &indexType,
		},
		cli.StringFlag{
			Name: "index-url",
			Value: "http://localhost:9200",
			Usage: "path to index",
			EnvVar: "APP_CRAWLER_INDEX",
			Destination: &indexUrl,
		},
		cli.StringFlag{
			Name: "token",
			Usage: "Socrata app token",
			EnvVar: "APP_SOCRSTA_TOKEN",
			Destination: &sfgovAppToken,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run the `SF Movies` crawler api",
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
