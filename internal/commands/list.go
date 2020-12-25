package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/list"
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"

	"github.com/spf13/cobra"
)

func listCommand(hosts goodhosts.Hosts) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the hosts in the host file",
		Run: func(cmd *cobra.Command, args []string) {
			list.List(hosts)
		},
	}

	return cmd

}
