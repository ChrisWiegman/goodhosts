package goodhosts

import (
	"fmt"
	"net"
	"strings"
)

// HostsLine Represents a single line in the hosts file.
type HostsLine struct {
	IP      string
	Hosts   []string
	Comment string
	Raw     string
	Err     error
}

// NewHostsLine Return a new instance of ```HostsLine```.
func NewHostsLine(raw string) HostsLine {
	fields := strings.Fields(raw)

	if len(fields) == 0 {
		return HostsLine{Raw: raw}
	}

	output := HostsLine{Raw: raw}

	if !IsComment(output.Raw) {
		rawIP := fields[0]
		if net.ParseIP(rawIP) == nil {
			output.Err = fmt.Errorf("bad hosts line: %q", raw)
		}

		output.IP = rawIP
		var outputFields []string

		for i, field := range fields {
			if IsComment(field) {
				output.Comment = strings.Join(fields[i:], " ")
				output.Comment = output.Comment[1:]
				break
			}
			outputFields = append(outputFields, field)
		}

		output.Hosts = outputFields[1:]
	}

	return output
}
