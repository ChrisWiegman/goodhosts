package list

import (
	"fmt"

	"github.com/ChrisWiegman/goodhosts/v4/internal/flags"
	"github.com/ChrisWiegman/goodhosts/v4/pkg/goodhosts"

	"github.com/spf13/cobra"
)

// List lists all the entries in the hosts file
func List(cmd *cobra.Command, args []string) error {

	hosts, err := goodhosts.NewHosts(flags.Section)
	if err != nil {
		return err
	}

	total := 0

	for _, line := range hosts.FileLines {

		var lineOutput string

		if line.Raw == "" {
			continue
		}

		if goodhosts.IsComment(line.Raw) && !flags.AllLines {
			continue
		}

		lineOutput = fmt.Sprintf("%s", line.Raw)
		if line.Err != nil {
			lineOutput = fmt.Sprintf("%s # <<< Malformated!", lineOutput)
		}

		if !goodhosts.IsComment(line.Raw) {
			total++
		}

		fmt.Println(lineOutput)
	}

	fmt.Println("") // Add a blank line
	fmt.Println(fmt.Sprintf("Total Host Entries: %d", total))

	return nil

}
