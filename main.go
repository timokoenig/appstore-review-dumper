package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "AppStore Review Dumper",
		Usage: "Dumps recent AppStore reviews of given app",
		Commands: []*cli.Command{
			{
				Name:    "dump",
				Aliases: []string{"d"},
				Usage:   "dump dump dumper",
				Action: func(c *cli.Context) error {
					dumpReviews(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all saved reviews from given URL",
				Action: func(c *cli.Context) error {
					listReviews(c.Args().Get(0))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
