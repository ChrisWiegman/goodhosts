package commands

import (
	"github.com/ChrisWiegman/goodhosts/v4/internal/remove"

	"github.com/spf13/cobra"
)

func removeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove <ip> <host>...",
		Short: "Removes an IP/Host combination from the hosts file",
		RunE:  remove.Remove,
		Args:  cobra.MinimumNArgs(2),
	}

	return cmd
}
