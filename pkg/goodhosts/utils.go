package goodhosts

import (
	"fmt"
	"strings"
)

// intemInSlice Returns true if the item string is in the provided slice or false
func itemInSlice(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

// buildRawLine builds a line for insertion into the hosts file
func buildRawLine(ip, host, comment string) string {
	output := ip

	output = fmt.Sprintf("%s %s", output, host)

	if len(comment) > 0 {
		comment = "#" + comment
		output = fmt.Sprintf("%s %s", output, comment)
	}

	return output

}

// IsComment Return ```true``` if the line is a comment.
func IsComment(line string) bool {

	trimLine := strings.TrimSpace(line)
	isComment := strings.HasPrefix(trimLine, commentChar)

	return isComment

}
