package main

import (
	"github.com/tiennampham23/berachain-airdrop/cmd/bot"
	"github.com/tiennampham23/berachain-airdrop/pkg/log"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	relativePathFlag = "relative-path"
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{},
				Action:  run,
				Flags:   flags(),
			},
		}}

	err := app.Run(os.Args)
	if err != nil {
		log.Logger().Fatalw("Failed to run this service", "err", err)
	}

}

func run(c *cli.Context) error {
	relativePath := c.String(relativePathFlag)

	b, err := bot.NewBot(relativePath)
	if err != nil {
		return err
	}
	return b.Start()
}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     relativePathFlag,
			EnvVars:  []string{"RELATIVE_PATH"},
			Required: true,
			Aliases:  []string{"f"},
		},
	}
}
