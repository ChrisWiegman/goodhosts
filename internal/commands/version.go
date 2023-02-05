package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version, Timestamp string

type VersionInfo struct {
	Version, Timestamp string
}

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Displays version information for the goodhosts.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Build Time: %s\n", Timestamp)
		},
		Args: cobra.NoArgs,
	}

	return cmd
}
