package check

import (
	"fmt"
	"os"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"
)

// Check checks the hosts file that the provided hosts are assigned to the ip
func Check(args []string, hosts goodhosts.Hosts) {

	hasErr := false

	ip := args[0]
	hostEntries := args[1:]

	for _, hostEntry := range hostEntries {

		if !hosts.Has(ip, hostEntry, false) {

			fmt.Fprintln(os.Stderr, fmt.Sprintf("%s %s is not in the hosts file", ip, hostEntry))
			hasErr = true

		}
	}

	if hasErr {
		os.Exit(1)
	}
}
