package commands

import (
	"os"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/flags"

	"github.com/spf13/cobra"
)

// Execute Runs the kana command.
func Execute() {

	cmd, err := rootCommand()
	if err != nil {
		os.Exit(1)
	}

	err = cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func rootCommand() (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:   "goodhosts [command]",
		Short: "Simple hosts file management.",
		Long:  "Goodhosts is a simple host file management utility.",
		Args:  cobra.MinimumNArgs(1),
	}

	flags.AddVerboseFlag(cmd)
	flags.AddSectionFlag(cmd)

	cmd.AddCommand(
		checkCommand(),
		listCommand(),
		addCommand(),
		removeCommand(),
		removeSectionCommand(),
	)

	return cmd, nil

}
