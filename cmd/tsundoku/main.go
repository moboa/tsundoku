package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tsundoku"
	app.Usage = "manga downloader"
	app.UsageText = app.Name + " [--help] [--version] [<url>]"
	app.Action = func(c *cli.Context) error {
		if c.NArg() != 1 {
			errorMsg := "Incorrect number of arguments. Please run " + app.Name + " -h."
			return cli.NewExitError(errorMsg, 1)
		}

		chapterURL, err := url.Parse(c.Args().Get(0))
		if err != nil {
			panic(err)
		}

		fmt.Println(chapterURL.Hostname())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
