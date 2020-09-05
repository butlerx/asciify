package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var version = "development"

const defaultWidth = 100

func main() {
	cwd, _ := os.Getwd()
	app := &cli.App{
		Name:      "asciify",
		Usage:     "asciffy all the things",
		UsageText: "asciffy [options] IMAGE",
		Description: `
		Every project needs a logo, so why not an ascii one.
		Convert image to ascii.
		Supports injecting image to readme`,
		Authors: []*cli.Author{
			{
				Name:  "Cian Butler",
				Email: "butlerx@notthe.cloud",
			},
		},
		Version:              version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "readme, r",
				Usage: "Path to readme.md to inject ascii too",
			},
			&cli.StringFlag{
				Name:  "path",
				Usage: "Path to output txt file too",
				Value: cwd,
			},
			&cli.BoolFlag{
				Name:  "print, p",
				Usage: "Print image to terminal",
			},
			&cli.IntFlag{
				Name:  "width, w",
				Usage: "Width of output",
				Value: defaultWidth,
			},
		},
		Action: run,
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
