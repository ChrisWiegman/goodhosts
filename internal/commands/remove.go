package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/remove"

	"github.com/spf13/cobra"
)

func removeCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "remove <ip> <host>...",
		Short: "Removes an IP/Host combination from the hosts file",
		RunE: remove.Remove,
	}

	return cmd

}
