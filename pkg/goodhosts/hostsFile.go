package goodhosts

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
)

const sectionStart = "### Start Hosts for"
const sectionEnd = "### End Hosts for"

// Hosts Represents a hosts file.
type Hosts struct {
	Path         string
	Section      string
	FileLines    []HostsLine
	SectionLines []HostsLine
}

var defaultFilePermissions = 660

// IsWritable Return ```true``` if hosts file is writable.
func (h *Hosts) IsWritable() bool {
	_, err := os.OpenFile(h.Path, os.O_WRONLY, os.FileMode(defaultFilePermissions))
	return err == nil
}

// Load the hosts file into ```l.Lines```.
// ```Load()``` is called by ```NewHosts()``` and ```Hosts.Flush()``` so you
// generally you won't need to call this yourself.
func (h *Hosts) Load() error {
	var fileLines []HostsLine
	var sectionLines []HostsLine
	var inSection bool

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

		if line.Raw == fmt.Sprintf("%s %s", sectionEnd, h.Section) {
			inSection = false
		}

		if inSection {
			sectionLines = append(sectionLines, line)
		}

		if line.Raw == fmt.Sprintf("%s %s", sectionStart, h.Section) {
			inSection = true
		}

		if !inSection && line.Raw != fmt.Sprintf("%s %s", sectionEnd, h.Section) {
			fileLines = append(fileLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	h.FileLines = fileLines
	h.SectionLines = sectionLines

	return nil
}

// Flush any changes made to hosts file.
func (h *Hosts) Flush() error {
	file, err := os.Create(h.Path)
	if err != nil {
		return err
	}

	if len(h.SectionLines) > 0 {
		if len(h.Section) > 0 {
			h.FileLines = append(h.FileLines, NewHostsLine(""), NewHostsLine(fmt.Sprintf("%s %s", sectionStart, h.Section)))
		}

		h.FileLines = append(h.FileLines, h.SectionLines...)

		if len(h.Section) > 0 {
			h.FileLines = append(h.FileLines, NewHostsLine(fmt.Sprintf("%s %s", sectionEnd, h.Section)), NewHostsLine(""))
		}
	}

	var isBlank bool
	w := bufio.NewWriter(file)

	for _, line := range h.FileLines {
		if !isBlank || len(line.Raw) > 1 {
			fmt.Fprintf(w, "%s%s", line.Raw, eol)
		}

		if len(line.Raw) < 2 {
			isBlank = true
		} else {
			isBlank = false
		}
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
		if h.Has(ip, host, true) {
			return fmt.Errorf("%s has already been assigned", host)
		}

		if !h.Has(ip, host, false) {
			endLine := NewHostsLine(buildRawLine(ip, host, comment))
			endLine.Comment = comment
			h.SectionLines = append(h.SectionLines, endLine)
		}
	}

	return nil
}

// Has Return a bool if ip/host combo in hosts file.
func (h *Hosts) Has(ip, host string, forceFile bool) bool {
	pos := h.getHostPosition(ip, host, forceFile)

	return pos != -1
}

// RemoveSection removes an entire section from the hostsfile
func (h *Hosts) RemoveSection() error {
	if h.Section == "" {
		return errors.New("no section provided")
	}

	h.SectionLines = nil

	return nil
}

// Remove an entry from the hosts file.
func (h *Hosts) Remove(ip string, hosts ...string) error {
	var outputLines []HostsLine
	inputLines := h.SectionLines

	if h.Section == "" {
		inputLines = h.FileLines
	}

	if net.ParseIP(ip) == nil {
		return fmt.Errorf("%q is an invalid IP address", ip)
	}

	for _, line := range inputLines {
		// Bad lines or comments just get re-added.
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

			if len(line.Comment) > 0 {
				newLineRaw = fmt.Sprintf("%s #%s", newLineRaw, line.Comment)
			}

			newLine := NewHostsLine(newLineRaw)
			outputLines = append(outputLines, newLine)
		}
	}

	if h.Section == "" {
		h.FileLines = outputLines
	} else {
		h.SectionLines = outputLines
	}

	return nil
}

func (h *Hosts) getHostPosition(ip, host string, forceFile bool) int {
	checkLines := h.FileLines

	if len(h.Section) > 0 && !forceFile {
		checkLines = h.SectionLines
	}

	for i := range checkLines {
		line := checkLines[i]
		if !IsComment(line.Raw) && line.Raw != "" {
			if ip == line.IP && itemInSlice(host, line.Hosts) {
				return i
			}
		}
	}

	return -1
}
