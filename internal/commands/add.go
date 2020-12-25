package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/add"
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"

	"github.com/spf13/cobra"
)

func addCommand(hosts goodhosts.Hosts) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add <ip> <host>...",
		Short: "Add an IP and host to the hosts file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return add.Add(args, hosts)
		},
		Args:  cobra.MinimumNArgs(2),
	}

	return cmd

}
