package config

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "Simple Backend Gorent With golang",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)

}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
