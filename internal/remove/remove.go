package remove

import (
	"errors"
	"fmt"

	"github.com/ChrisWiegman/goodhosts/v4/internal/flags"
	"github.com/ChrisWiegman/goodhosts/v4/pkg/goodhosts"

	"github.com/spf13/cobra"
)

// Remove deletes a host/ip entry.
func Remove(cmd *cobra.Command, args []string) error {
	hosts, err := goodhosts.NewHosts(flags.Section)
	if err != nil {
		return err
	}

	ip := args[0]
	hostEntries := args[1:]

	if !hosts.IsWritable() {
		return errors.New("host file not writable. Try running with elevated privileges")
	}

	err = hosts.Remove(ip, hostEntries...)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return hosts.Flush()
}
