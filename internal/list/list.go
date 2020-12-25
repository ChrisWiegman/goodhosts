package list

import (
	"fmt"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/flags"
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"
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

		if goodhosts.IsComment(line.Raw) && !flags.AllLines {
			continue
		}

		lineOutput = fmt.Sprintf("%s", line.Raw)
		if line.Err != nil {
			lineOutput = fmt.Sprintf("%s # <<< Malformated!", lineOutput)
		}
		total++

		fmt.Println(lineOutput)
	}

	fmt.Printf("\nTotal: %d\n", total)

	return nil

}
