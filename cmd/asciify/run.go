package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/butlerx/ascii-logo/internal/pkg/asciify"
	"github.com/butlerx/ascii-logo/pkg/ascii"
	"github.com/urfave/cli"
)

func run(ctx *cli.Context) error {
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
