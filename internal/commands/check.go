package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/check"
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"

	"github.com/spf13/cobra"
)

func checkCommand(hosts goodhosts.Hosts) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "check <ip> <host>...",
		Short: "Check if an ip and host(s) is present in the hosts file",
		Run: func(cmd *cobra.Command, args []string) {
			check.Check(args, hosts)
		},
		Args:  cobra.MinimumNArgs(2),
	}

	return cmd

}
