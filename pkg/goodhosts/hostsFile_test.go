package goodhosts

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	err := hosts.Add("10.0.0.7", "", "brada")
	assert.NoError(t, err)

	err = hosts.Add("127.0.0.1", "", "yadda")
	assert.Error(t, err)

	hosts.FileLines = append(hosts.FileLines, hosts.SectionLines...)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 brada"),
	}

	assert.Equal(t, hosts.FileLines, expectedLines)
}

func TestHostsAddWithComment(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
	}

	err := hosts.Add("10.0.0.7", "Test Comment", "brada", "yadda")
	assert.NoError(t, err)

	hosts.FileLines = append(hosts.FileLines, hosts.SectionLines...)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada"),
		NewHostsLine("10.0.0.7 brada #Test Comment"),
		NewHostsLine("10.0.0.7 yadda #Test Comment"),
	}

	assert.Equal(t, hosts.FileLines, expectedLines)
}

func TestHostsAddWhenIpDoesntExist(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
	}

	err := hosts.Add("10.0.0.7", "", "brada", "yadda")
	assert.NoError(t, err)

	hosts.FileLines = append(hosts.FileLines, hosts.SectionLines...)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
		NewHostsLine("10.0.0.7 yadda"),
	}

	assert.Equal(t, hosts.FileLines, expectedLines)
}

func TestHostsRemoveWhenLastHostIpCombo(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	err := hosts.Remove("10.0.0.7", "nada")
	assert.NoError(t, err)

	expectedLines := []HostsLine{NewHostsLine("127.0.0.1 yadda")}

	assert.Equal(t, hosts.FileLines, expectedLines)
}

func TestHostsRemoveWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)

	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 nada brada")}

	err := hosts.Remove("10.0.0.7", "nada")
	assert.NoError(t, err)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada")}

	assert.Equal(t, hosts.FileLines, expectedLines)
}

func TestHostsRemoveMultipleEntries(t *testing.T) {
	hosts := new(Hosts)
	hosts.FileLines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda nadda prada")}

	err := hosts.Remove("127.0.0.1", "yadda", "prada")
	assert.NoError(t, err)

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

	err := hosts.Remove("10.0.0.7", "nada")
	assert.NoError(t, err)

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"),
		NewHostsLine("10.0.0.7 brada"),
	}

	assert.Equal(t, hosts.FileLines, expectedLines)
}
