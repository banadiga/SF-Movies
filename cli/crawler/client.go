package main

import (
	"../../api"
	. "../../api/crawler"
	"../../crawler-client"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var host string

func start(context *cli.Context) {
	log.Println("Getting crawler information...")

	status, err := crawler.NewCliente(host).Start()
	if err != nil {
		printError(err)
		return
	}

	printStatus(status);
}

func status(context *cli.Context) {
	log.Println("Getting crawler information...")

	status, err := crawler.NewCliente(host).Status()
	if err != nil {
		printError(err)
		return
	}

	printStatus(status);
}

func stop(context *cli.Context) {
	log.Println("Getting crawler information...")

	status, err := crawler.NewCliente(host).Stop()
	if err != nil {
		printError(err)
		return
	}

	printStatus(status);
}

func printError(err error) {
	log.Fatalf("Error: %s", err.Error())
}

func printStatus(status *Status) {
	log.Println("CrawlerStatus:")
	log.Printf("Active: %t", status.Active)
	log.Printf("Offset: %d", status.Offset)
}

func main() {
	app := cli.NewApp()
	app.Name = "SF Movies crawler cli"
	app.Usage = "cli to start/stop the SF Movies crawler and get status of the crawler."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "http://localhost:" + api.Port,
			Usage:  "SF Movies crawler host",
			EnvVar: "CRAWLER_API_HOST",
			Destination: &host,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the crawler",
			Action: start,
		},
		{
			Name:  "status",
			Usage: "Status of the crawler",
			Action: status,
		},
		{
			Name:  "stop",
			Usage: "Stop the crawler",
			Action: stop,
		},
	}

	app.Run(os.Args)
}
