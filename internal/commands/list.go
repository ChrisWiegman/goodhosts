package commands

import (
	"github.com/ChrisWiegman/goodhosts/v4/internal/flags"
	"github.com/ChrisWiegman/goodhosts/v4/internal/list"

	"github.com/spf13/cobra"
)

func listCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [-all]",
		Short: "List all the hosts in the host file",
		RunE:  list.List,
	}

	flags.AddAllLinesFlag(cmd)

	return cmd
}
