package add

import (
	"errors"
	"fmt"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/pkg/goodhosts"
)

// Add checks the hosts file that the provided hosts are assigned to the ip
func Add(args []string, hosts goodhosts.Hosts) error {

	ip := args[0]
	hostEntries := args[1:]

	if !hosts.IsWritable() {
		return errors.New("Host file not writable. Try running with elevated privileges.")
	}

	err := hosts.Add(ip, "Test Comment", hostEntries...)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return hosts.Flush()

}
