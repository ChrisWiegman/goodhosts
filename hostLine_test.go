package goodhosts

import (
	"fmt"
	"testing"
)

func TestHostsLineIsComment(t *testing.T) {
	comment := "   # This is a comment   "
	line := NewHostsLine(comment)
	result := IsComment(line.Raw)
	if !result {
		t.Error(fmt.Sprintf("'%s' should be a comment", comment))
	}
}

func TestNewHostsLineWithEmptyLine(t *testing.T) {
	line := NewHostsLine("")
	if line.Raw != "" {
		t.Error("Failed to load empty line.")
	}
}
