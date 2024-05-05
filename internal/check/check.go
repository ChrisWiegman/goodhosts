package check

import (
	"fmt"
	"os"

	"github.com/ChrisWiegman/goodhosts/v4/internal/flags"
	"github.com/ChrisWiegman/goodhosts/v4/pkg/goodhosts"

	"github.com/spf13/cobra"
)

// Check checks the hosts file that the provided hosts are assigned to the ip.
func Check(cmd *cobra.Command, args []string) error {
	hosts, err := goodhosts.NewHosts(flags.Section)
	if err != nil {
		return err
	}

	hasErr := false

	ip := args[0]
	hostEntries := args[1:]

	for _, hostEntry := range hostEntries {
		if !hosts.Has(ip, hostEntry, false) {
			fmt.Fprintf(os.Stderr, "%s %s is not in the hosts file", ip, hostEntry)
			hasErr = true
		}
	}

	if hasErr {
		os.Exit(1)
	}

	return nil
}
