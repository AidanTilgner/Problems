package problems

import (
	"github.com/urfave/cli/v2"
)

func Problem() *cli.Command {
	cmd := &cli.Command{
		Name:  "problem",
		Usage: "Manage problems.",
		Subcommands: []*cli.Command{
			SolveProblem(),
		},
	}

	return cmd
}

func SolveProblem() *cli.Command {
	cmd := &cli.Command{
		Name:  "solve",
		Usage: "Solve a problem.",
		Action: func(cCtx *cli.Context) error {
			SolveProblem()
			return nil
		},
	}

	return cmd
}
