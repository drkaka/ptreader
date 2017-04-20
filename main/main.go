package main

import (
	"fmt"
	"os"

	"github.com/drkaka/ptreader/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ptreader"
	app.Usage = "command line application to read db of procrastitracker"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Specify the input db file.",
		},
	}

	app.Before = func(c *cli.Context) error {
		dbPath := c.GlobalString("f")
		if st, err := os.Stat(dbPath); err != nil {
			return err
		} else if st.IsDir() {
			return fmt.Errorf("%s is not a file", dbPath)
		}
		return nil
	}

	app.Commands = []cli.Command{
		cmd.NewListCMD(),
	}

	app.Run(os.Args)
}
