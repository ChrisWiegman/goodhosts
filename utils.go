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

func buildRawLine(ip string, host, comment string) string {
	output := ip

	if len(comment) > 0 {
		comment = "#" + comment
	}
	output = fmt.Sprintf("%s %s %s", output, host, comment)

	return output
}

// IsComment Return ```true``` if the line is a comment.
func IsComment(line string) bool {
	trimLine := strings.TrimSpace(line)
	isComment := strings.HasPrefix(trimLine, commentChar)
	return isComment
}
