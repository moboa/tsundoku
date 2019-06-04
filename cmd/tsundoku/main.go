package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/moboa/tsundoku/manga"
	"github.com/moboa/tsundoku/save"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tsundoku"
	app.Version = "0.0.1"
	app.Usage = "manga downloader"
	app.UsageText = app.Name + " [--help/-h] [--version/-v] [--output/-o] [<url>]"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: ".",
			Usage: "Save files to output `DIR`",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() != 1 {
			fmt.Println(c.NArg())
			errorMsg := "Incorrect number of arguments. Please run " + app.Name + " -h."
			return cli.NewExitError(errorMsg, 1)
		}

		pageURL, err := url.Parse(c.Args().Get(0))
		if err != nil {
			panic(err)
		}

		images := manga.FetchPageImages(pageURL)
		save.ToFiles(images, c.String("output"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
