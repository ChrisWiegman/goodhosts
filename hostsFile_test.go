package goodhosts

import (
	"reflect"
	"testing"
)

func TestHostsHas(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
	}

	// We should find this entry.
	if !hosts.Has("10.0.0.7", "nada", false) {
		t.Error("Couldn't find entry in hosts file.")
	}

	// We shouldn't find this entry
	if hosts.Has("10.0.0.7", "shuda", false) {
		t.Error("Found entry that isn't in hosts file.")
	}
}

func TestHostsHasDoesntFindMissingEntry(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	if hosts.Has("10.0.0.7", "brada", false) {
		t.Error("Found missing entry.")
	}
}

func TestHostsAddWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
	}

	hosts.Add("10.0.0.7", "", "brada")
	hosts.Add("127.0.0.1", "", "yadda")

	checkLines := append(hosts.FileLines, hosts.SectionLines...)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 brada"),
	}

	if !reflect.DeepEqual(checkLines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsAddWithComment(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
	}

	hosts.Add("10.0.0.7", "Test Comment", "brada", "yadda")

	checkLines := append(hosts.FileLines, hosts.SectionLines...)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 brada #Test Comment"),
		NewHostsLine("10.0.0.7 yadda #Test Comment"),
	}

	if !reflect.DeepEqual(checkLines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsAddWhenIpDoesntExist(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
	}

	hosts.Add("10.0.0.7", "", "brada", "yadda")

	checkLines := append(hosts.FileLines, hosts.SectionLines...)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	if !reflect.DeepEqual(checkLines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsRemoveWhenLastHostIpCombo(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	hosts.Remove("10.0.0.7", "nada")

	expectedLines := []HostsLine{NewHostsLine("127.0.0.1 yadda")}

	if !reflect.DeepEqual(hosts.FileLines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}

func TestHostsRemoveWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)

	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada brada")}

	hosts.Remove("10.0.0.7", "nada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada")}

	if !reflect.DeepEqual(hosts.FileLines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}

func TestHostsRemoveMultipleEntries(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda nadda prada")}

	hosts.Remove("127.0.0.1", "yadda", "prada")
	if hosts.FileLines[0].Raw != "127.0.0.1 nadda" {
		t.Error("Failed to remove multiple entries.")
	}
}

func TestHostsRemoveLineWithComments(t *testing.T) {
	hosts := new(Hosts)

	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
	}

	nadaLine := NewHostsLine("10.0.0.7 nada")
	nadaLine.Comment = "Test comment"

	hosts.FileLines = append(hosts.FileLines, nadaLine)

	hosts.Remove("10.0.0.7", "nada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
	}

	if !reflect.DeepEqual(hosts.FileLines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}
