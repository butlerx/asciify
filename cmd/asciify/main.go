package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "asciify"
	app.Usage = "asciffy all the things"
	app.UsageText = "asciffy [options] IMAGE"
	app.Description = `
		Every project needs a logo, so why not an ascii one.
		Convert image to ascii.
		Supports injecting image to readme`
	app.Authors = []cli.Author{
		{
			Name:  "Cian Butler",
			Email: "butlerx@notthe.cloud",
		},
	}
	app.Version = "1.0.0"
	app.EnableBashCompletion = true
	cwd, _ := os.Getwd()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "readme, r",
			Usage: "Path to readme.md to inject ascii too",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "Path to output txt file too",
			Value: cwd,
		},
		cli.BoolFlag{
			Name:  "print, p",
			Usage: "Print image to terminal",
		},
		cli.IntFlag{
			Name:  "width, w",
			Usage: "Width of output",
			Value: 100,
		},
	}
	app.Action = run
	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
