package main

import (
	"fmt"
	"os"

	"gitea.chriswiegman.com/ChrisWiegman/goodhosts"
	"github.com/docopt/docopt-go"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	usage := `Goodhosts - simple hosts file management.

Usage:
  goodhosts check <ip> <host>...
  goodhosts add <ip> <host>...
  goodhosts (rm|remove) <ip> <host>...
  goodhosts list [--all]
  goodhosts removesection <section>
  goodhosts -h | --help
  goodhosts --version

Options:
  --all         Display comments when listing.
  -h --help     Show this screen.
  --version     Show the version.`

	args, _ := docopt.Parse(usage, nil, true, "Goodhosts 3.2.0", false)

	hosts, err := goodhosts.NewHosts("")
	check(err)

	if args["removesection"].(bool) {

		section := args["<section>"].(string)

		hosts, err := goodhosts.NewHosts(section)
		check(err)

		if !hosts.IsWritable() {
			fmt.Fprintln(os.Stderr, "Host file not writable. Try running with elevated privileges.")
			os.Exit(1)
		}

		err = hosts.RemoveSection()
		if err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err.Error()))
			os.Exit(2)
		}

		err = hosts.Flush()
		check(err)

		return
	}
}
