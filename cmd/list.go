package cmd

import (
	"fmt"

	"github.com/drkaka/ptreader"
	"github.com/urfave/cli"
)

// NewListCMD to list information.
func NewListCMD() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "list information",
		Action: listAction,
	}
}

func listAction(c *cli.Context) error {
	f := c.GlobalString("f")
	n, err := ptreader.Read(f)
	if err != nil {
		return err
	}

	fmt.Println(n.Name)
	for i := 0; i < len(n.SubNodes); i++ {
		one := n.SubNodes[i]
		seconds := ptreader.GetNodeTime(&one)
		fmt.Printf("%d: %s, %s \n", i, one.Name, ptreader.FormatSeconds(seconds))
	}
	return nil
}
