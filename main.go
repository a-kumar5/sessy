package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/a-kumar5/sessy/utils"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "sessy",
		Usage: "smart session manager for devops engineer",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list all items in current directory",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					output, err := utils.ListFiles()
					if err != nil {
						return err
					}
					fmt.Printf(string(output))
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
