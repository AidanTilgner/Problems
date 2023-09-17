package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"problem-solver/problems"
)

func main() {
	app := &cli.App{
		Name:  "Problem CLI.",
		Usage: "Deal with problems.",
		Commands: []*cli.Command{
			problems.Problem(),
		},
	}

	app.Run(os.Args)
}
