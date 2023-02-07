package cmd

import (
	"github.com/urfave/cli/v2"

	"main/access"
)

func Commands() *cli.App {
	app := &cli.App{
		Name:        "kloner",
		Description: "simple cli tool to access linux servers through local cli",
		Commands: []*cli.Command{
			{
				Name:    "access",
				Aliases: []string{"a"},
				Usage:   "add ssh access details",
				Action: func(c *cli.Context) error {
					access.ConnectToServerWithPrivatePublicKeys(
						c.String("user"),
						c.String("host"),
						c.String("port"),
						c.String("pKey"),
						c.String("type"),
					)
					return nil
				},
				ArgsUsage:   ` `,
				Description: `Access linux server through ssh`,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "pKey",
						Usage: "path to your private key",
					},
					&cli.StringFlag{
						Name:  "user",
						Usage: "the username to the linux server",
					},
					&cli.StringFlag{
						Name:  "host",
						Usage: "your server ip or domain",
					},
					&cli.StringFlag{
						Name:  "port",
						Usage: "your access port or just 22 if no custom port",
					},
					&cli.StringFlag{
						Name:  "type",
						Usage: "private key type e.g. pem, putty or rsa",
					},
				},
			},
		},
	}

	return app
}