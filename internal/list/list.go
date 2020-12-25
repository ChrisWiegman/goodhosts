package list

import (
	"fmt"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"
)

// List lists all the entries in the hosts file
func List(hosts goodhosts.Hosts) {

	total := 0
		for _, line := range hosts.FileLines {
			var lineOutput string

			if goodhosts.IsComment(line.Raw) {
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

		return
}
