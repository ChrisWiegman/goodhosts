package flags

import "github.com/spf13/cobra"

// Comment Adds a comment to the entry for better identification later
var Comment string

// Section Is a named section with which to edit hosts in.
var Section string

// Verbose Set to true for verbose output on a given command.
var Verbose bool

// AddCommentFlag Adds the comment flag to a cobra command.
func AddCommentFlag(command *cobra.Command) {
	command.Flags().StringVar(&Comment, "comment", "", "Add a comment to the entry for better identification later.")
}

// AddSectionFlag Adds the section flag to a cobra command.
func AddSectionFlag(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&Section, "section", "s", "", "A named section with which to edit hosts in.")
}

// AddVerboseFlag Adds the verbose flag to a cobra command.
func AddVerboseFlag(command *cobra.Command) {
	command.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose (detailed) output.")
}
