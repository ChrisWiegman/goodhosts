package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/v4/internal/add"
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/v4/internal/flags"

	"github.com/spf13/cobra"
)

func addCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add <ip> <host>...",
		Short: "Add an IP and host to the hosts file",
		RunE: add.Add,
		Args:  cobra.MinimumNArgs(2),
	}

	flags.AddCommentFlag(cmd)

	return cmd

}
