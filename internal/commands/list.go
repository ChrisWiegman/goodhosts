package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/list"

	"github.com/spf13/cobra"
)

func listCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the hosts in the host file",
		RunE: list.List,
	}

	return cmd

}
