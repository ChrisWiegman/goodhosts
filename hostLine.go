package goodhosts

import (
	"fmt"
	"net"
	"strings"
)

// HostsLine Represents a single line in the hosts file.
type HostsLine struct {
	IP    string
	Hosts []string
	Raw   string
	Err   error
}

// IsComment Return ```true``` if the line is a comment.
func (l HostsLine) IsComment() bool {
	trimLine := strings.TrimSpace(l.Raw)
	isComment := strings.HasPrefix(trimLine, commentChar)
	return isComment
}

// NewHostsLine Return a new instance of ```HostsLine```.
func NewHostsLine(raw string) HostsLine {

	fields := strings.Fields(raw)
	if len(fields) == 0 {
		return HostsLine{Raw: raw}
	}

	output := HostsLine{Raw: raw}
	if !output.IsComment() {
		rawIP := fields[0]
		if net.ParseIP(rawIP) == nil {
			output.Err = fmt.Errorf("Bad hosts line: %q", raw)
		}

		output.IP = rawIP
		output.Hosts = fields[1:]
	}

	return output
}
