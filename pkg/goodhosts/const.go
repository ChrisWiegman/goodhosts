//go:build linux || darwin

package goodhosts

const hostsFilePath = "/etc/hosts"
const eol = "\n"
const commentChar string = "#"
