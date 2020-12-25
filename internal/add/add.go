package add

import (
	"errors"
	"fmt"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/flags"
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"

	"github.com/spf13/cobra"
)

// Add checks the hosts file that the provided hosts are assigned to the ip
func Add(cmd *cobra.Command, args []string) error {

	hosts, err := goodhosts.NewHosts(flags.Section)
	if err != nil {
		return err
	}

	ip := args[0]
	hostEntries := args[1:]

	if !hosts.IsWritable() {
		return errors.New("host file not writable. Try running with elevated privileges")
	}

	err = hosts.Add(ip, flags.Comment, hostEntries...)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return hosts.Flush()

}
