package goodhosts

import (
	"reflect"
	"testing"
)

func TestHostsHas(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
	}

	// We should find this entry.
	if !hosts.Has("10.0.0.7", "nada") {
		t.Error("Couldn't find entry in hosts file.")
	}

	// We shouldn't find this entry
	if hosts.Has("10.0.0.7", "shuda") {
		t.Error("Found entry that isn't in hosts file.")
	}
}

func TestHostsHasDoesntFindMissingEntry(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	if hosts.Has("10.0.0.7", "brada") {
		t.Error("Found missing entry.")
	}
}

func TestHostsAddWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	hosts.Add("10.0.0.7", "", "brada", "yadda")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 yadda"),
		NewHostsLine("10.0.0.7 brada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsAddWithComment(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	hosts.Add("10.0.0.7", "Test Comment", "brada", "yadda")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	bradaLine := NewHostsLine("10.0.0.7 brada")
	yaddaLine := NewHostsLine("10.0.0.7 yadda")
	bradaLine.comment = "Test Comment"
	yaddaLine.comment = "Test Comment"

	expectedLines = append(expectedLines, bradaLine)
	expectedLines = append(expectedLines, yaddaLine)

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsAddWhenIpDoesntExist(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
	}

	hosts.Add("10.0.0.7", "", "brada", "yadda")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsRemoveWhenLastHostIpCombo(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	hosts.Remove("10.0.0.7", "nada")

	expectedLines := []HostsLine{NewHostsLine("127.0.0.1 yadda")}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}

func TestHostsRemoveWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)

	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada brada")}

	hosts.Remove("10.0.0.7", "nada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada")}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}

func TestHostsRemoveMultipleEntries(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda nadda prada")}

	hosts.Remove("127.0.0.1", "yadda", "prada")
	if hosts.Lines[0].Raw != "127.0.0.1 nadda" {
		t.Error("Failed to remove multiple entries.")
	}
}

func TestHostsRemoveLineWithComments(t *testing.T) {
	hosts := new(Hosts)

	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
	}

	nadaLine := NewHostsLine("10.0.0.7 nada")
	nadaLine.comment = "Test comment"

	hosts.Lines = append(hosts.Lines, nadaLine)

	hosts.Remove("10.0.0.7", "nada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
	}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}
