package problems

import (
	"fmt"
	brain "problem-solver/brain"
)


func Problem() *cli.Command {
	cmd := &cli.Command{
		Name:        "problem",
		Usage:       "Manage problems.",
		Subcommands: []*cli.Command{
			NewProblem(),
		}
	}

	return cmd
}

func SolveProblem() *cli.Command {
	cmd := &cli.Command{
		Name:        "solve",
		Usage:       "Solve a problem.",
		Action:	     func(cCtx *cli.Context) error {
			SolveProblem()
			return nil
		}
	}

	return cmd
}
