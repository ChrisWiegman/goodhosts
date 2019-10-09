package goodhosts

import "fmt"

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
