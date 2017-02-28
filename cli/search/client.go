package main

import (
	"../../api"
	. "../../search-client"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var host string

func actionSearch(c *cli.Context) {
	log.Println("Searching...")
	name := c.Args().Get(0)

	films, err := NewClient(host).Search(name)
	if err != nil {
		printError(err)
		return
	}
	printFilms(films)
}

func printError(err error) {
	log.Fatalf("Error: %s", err)
}

func printFilms(films *api.Films) {
	log.Println("Films:")
	log.Printf("Size: %d", films.Len())
}

func main() {
	app := cli.NewApp()
	app.Name = "SF Movies search cli"
	app.Usage = "cli to work with the SF Movies search api"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "http://localhost:" + api.Port,
			Usage:  "SF Movies service host",
			EnvVar: "SEARCH_API_HOST",
			Destination: &host,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "search",
			Usage: "(name) search, health check",
			Action: actionSearch,
			ArgsUsage: "[name]",
		},
	}

	app.Run(os.Args)
}
