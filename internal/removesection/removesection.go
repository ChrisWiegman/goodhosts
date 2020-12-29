package removesection

import (
	"errors"
	"fmt"

	"github.com/ChrisWiegman/goodhosts/v4/internal/flags"
	"github.com/ChrisWiegman/goodhosts/v4/pkg/goodhosts"

	"github.com/spf13/cobra"
)

// RemoveSection deletes a section from the hosts file
func RemoveSection(cmd *cobra.Command, args []string) error {

	if flags.Section == "" {
		return errors.New("you must provide the `--section` flag to use this command")
	}

	hosts, err := goodhosts.NewHosts(flags.Section)
	if err != nil {
		return err
	}

	if !hosts.IsWritable() {
		return errors.New("host file not writable. Try running with elevated privileges")
	}

	err = hosts.RemoveSection()
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return hosts.Flush()

}
