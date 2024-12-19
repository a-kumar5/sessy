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
			{
				Name:    "yaml",
				Aliases: []string{"yl"},
				Usage:   "create a yaml/yml file if name given it will create with that name or else default",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fileName := "default"
					if cmd.NArg() > 0 {
						fileName = cmd.Args().Get(0)
					}
					_, err := utils.CreateYaml(fileName)
					if err != nil {
						return err
					}
					fmt.Printf("%s.yml file created\n", fileName)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
