package goodhosts

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Hosts Represents a hosts file.
type Hosts struct {
	Path  string
	Lines []HostsLine
}

// IsWritable Return ```true``` if hosts file is writable.
func (h *Hosts) IsWritable() bool {
	_, err := os.OpenFile(h.Path, os.O_WRONLY, 0660)
	if err != nil {
		return false
	}

	return true
}

// Load the hosts file into ```l.Lines```.
// ```Load()``` is called by ```NewHosts()``` and ```Hosts.Flush()``` so you
// generally you won't need to call this yourself.
func (h *Hosts) Load() error {
	var lines []HostsLine

	file, err := os.Open(h.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := NewHostsLine(scanner.Text())
		if err != nil {
			return err
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	h.Lines = lines

	return nil
}

// Flush any changes made to hosts file.
func (h Hosts) Flush() error {
	file, err := os.Create(h.Path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)

	for _, line := range h.Lines {
		var comment string
		if len(line.comment) > 0 {
			comment = " #" + line.comment
		}
		fmt.Fprintf(w, "%s%s%s", line.Raw, comment, eol)
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return h.Load()
}

// Add an entry to the hosts file.
func (h *Hosts) Add(ip, comment string, hosts ...string) error {

	if net.ParseIP(ip) == nil {
		return fmt.Errorf("%q is an invalid IP address", ip)
	}

	for _, host := range hosts {

		endLine := NewHostsLine(buildRawLine(ip, host))
		endLine.comment = comment
		h.Lines = append(h.Lines, endLine)
	}

	return nil
}

// Has Return a bool if ip/host combo in hosts file.
func (h Hosts) Has(ip string, host string) bool {
	pos := h.getHostPosition(ip, host)

	return pos != -1
}

// Remove an entry from the hosts file.
func (h *Hosts) Remove(ip string, hosts ...string) error {
	var outputLines []HostsLine

	if net.ParseIP(ip) == nil {
		return fmt.Errorf("%q is an invalid IP address", ip)
	}

	for _, line := range h.Lines {

		// Bad lines or comments just get readded.
		if line.Err != nil || IsComment(line.Raw) || line.IP != ip {
			outputLines = append(outputLines, line)
			continue
		}

		var newHosts []string
		for _, checkHost := range line.Hosts {
			if !itemInSlice(checkHost, hosts) {
				newHosts = append(newHosts, checkHost)
			}
		}

		// If hosts is empty, skip the line completely.
		if len(newHosts) > 0 {
			newLineRaw := line.IP

			for _, host := range newHosts {
				newLineRaw = fmt.Sprintf("%s %s", newLineRaw, host)
			}
			newLine := NewHostsLine(newLineRaw)
			outputLines = append(outputLines, newLine)
		}
	}

	h.Lines = outputLines
	return nil
}

func (h Hosts) getHostPosition(ip string, host string) int {

	for i := range h.Lines {
		line := h.Lines[i]
		if !IsComment(line.Raw) && line.Raw != "" {
			if ip == line.IP && itemInSlice(host, line.Hosts) {
				return i
			}
		}
	}

	return -1
}

func (h Hosts) getIPPosition(ip string) int {
	for i := range h.Lines {
		line := h.Lines[i]
		if !IsComment(line.Raw) && line.Raw != "" {
			if line.IP == ip {
				return i
			}
		}
	}

	return -1
}