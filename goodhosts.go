package goodhosts

import (
	"os"
	"path/filepath"
)

// NewHosts Return a new instance of ``Hosts``.
func NewHosts() (Hosts, error) {
	osHostsFilePath := ""

	if os.Getenv("HOSTS_PATH") == "" {
		osHostsFilePath = os.ExpandEnv(filepath.FromSlash(hostsFilePath))
	} else {
		osHostsFilePath = os.Getenv("HOSTS_PATH")
	}

	hosts := Hosts{Path: osHostsFilePath}

	err := hosts.Load()
	if err != nil {
		return hosts, err
	}

	return hosts, nil
}
