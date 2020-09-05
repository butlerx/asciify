package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/butlerx/ascii-logo/internal/pkg/asciify"
	"github.com/butlerx/ascii-logo/pkg/ascii"
	"github.com/urfave/cli/v2"
)

var noFileError = errors.New("No file supplied, supply path or url of an image to convert it too ascii")

func run(ctx *cli.Context) error {
	if ctx.NArg() == 0 {
		return noFileError
	}

	imgPath := ctx.Args().First()

	file, err := asciify.GetFile(imgPath)
	if err != nil {
		return err
	}

	img, err := ascii.Generate(file, ctx.Int("width"))
	if err != nil {
		return err
	}

	if ctx.Bool("print") {
		fmt.Println(img)

		return nil
	} else if ctx.String("readme") != "" {
		return asciify.InjectReadMe(ctx.String("readme"), img)
	}

	txt := strings.TrimSuffix(
		filepath.Base(imgPath),
		filepath.Ext(filepath.Base(imgPath)),
	) + ".txt"
	if err := asciify.Save(filepath.Join(ctx.String("path"), txt), img); err != nil {
		return err
	}

	fmt.Println(filepath.Base(imgPath), "saved too", ctx.String("path"), "as", filepath.Base(txt))

	return nil
}
