package commands

import (
	"gitea.chriswiegman.com/ChrisWiegman/goodhosts/internal/removesection"

	"github.com/spf13/cobra"
)

func removeSectionCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "removesection --section=<section>",
		Short: "Removes an entire section from the hosts file",
		RunE: removesection.RemoveSection,
	}

	return cmd

}