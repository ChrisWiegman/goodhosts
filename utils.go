package goodhosts

import (
	"fmt"
	"strings"
)

func itemInSlice(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func buildRawLine(ip string, host string) string {
	output := ip
	output = fmt.Sprintf("%s %s", output, host)

	return output
}

// IsComment Return ```true``` if the line is a comment.
func IsComment(line string) bool {
	trimLine := strings.TrimSpace(line)
	isComment := strings.HasPrefix(trimLine, commentChar)
	return isComment
}
