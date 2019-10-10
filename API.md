# goodhosts
--
    import "github.com/lextoumbourou/goodhosts"


## Usage

#### type Hosts

```go
type Hosts struct {
	Path         string
	Section      string
	FileLines    []HostsLine
	SectionLines []HostsLine
}
```

Represents a hosts file.

#### func  NewHosts

```go
func NewHosts("") (Hosts, error)
```
Return a new instance of ``Hosts``.

#### func IsComment

```go
func IsComment(line string) bool
```
Return ```true``` if the string is a comment.

#### func (*Hosts) Add

```go
func (h *Hosts) Add(ip, comment string, hosts ...string) error
```
Add an entry to the hosts file.

#### func (Hosts) Flush

```go
func (h Hosts) Flush() error
```
Flush any changes made to hosts file.

#### func (Hosts) Has

```go
func (h Hosts) Has(ip string, host string, forceFile bool) bool
```
Return a bool if ip/host combo in hosts file.

#### func (*Hosts) IsWritable

```go
func (h *Hosts) IsWritable() bool
```
Return ```true``` if hosts file is writable.

#### func (*Hosts) Load

```go
func (h *Hosts) Load() error
```
Load the hosts file into ```l.Lines```. ```Load()``` is called by
```NewHosts()``` and ```Hosts.Flush()``` so you generally you won't need to call
this yourself.

#### func (*Hosts) Remove

```go
func (h *Hosts) Remove(ip string, hosts ...string) error
```
Remove an entry from the hosts file.

#### type HostsLine

```go
type HostsLine struct {
	IP      string
	Hosts   []string
	Comment string
	Raw     string
	Err     error
}
```

Represents a single line in the hosts file.

#### func  NewHostsLine

```go
func NewHostsLine(raw string) HostsLine
```
Return a new instance of ```HostsLine```.
