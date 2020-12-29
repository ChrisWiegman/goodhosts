package commands

import (
	"github.com/ChrisWiegman/goodhosts/v4/internal/check"

	"github.com/spf13/cobra"
)

func checkCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "check <ip> <host>...",
		Short: "Check if an ip and host(s) is present in the hosts file",
		RunE: check.Check,
		Args:  cobra.MinimumNArgs(2),
	}

	return cmd

}
