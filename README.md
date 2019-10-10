# Goodhosts

[![Build Status](https://travis-ci.org/ChrisWiegman/goodhosts.svg?branch=master)](https://travis-ci.org/ChrisWiegman/goodhosts)

This is a fork of [https://github.com/lextoumbourou/goodhosts] as we needed a bit deeper management.

Simple [hosts file](http://en.wikipedia.org/wiki/Hosts_%28file%29) (```/etc/hosts```) management in Go (golang).

## Features

* List, add, remove and check hosts file entries from code or the command-line.
* Windows support.

## Command-Line Usage

### List entries

```bash
$ goodhosts list
127.0.0.1 localhost
10.0.0.5 my-home-server xbmc-server
10.0.0.6 my-desktop
```

Add ```--all``` flag to include comments.

### Check for an entry

```bash
$ goodhosts check 127.0.0.1 facebook.com
```

### Add an entry

```bash
$ goodhosts add 127.0.0.1 facebook.com
```

Or *entries*.

```bash
$ goodhosts add 127.0.0.1 facebook.com twitter.com gmail.com
```

### Remove an entry

```bash
$ goodhosts rm 127.0.0.1 facebook.com
```

Or *entries*.

```bash
$ goodhosts rm 127.0.0.1 facebook.com twitter.com gmail.com
```

### More

```bash
$ goodhosts --help
```

## API Usage

### Installation

```bash
$ go get github.com/ChrisWiegman/goodhosts
```

### List entries

```go
package main

import (
    "fmt"
    "github.com/ChrisWiegman/goodhosts"
)

func main() {
    hosts := goodhosts.NewHosts()

    for _, line := range hosts.Lines {
        fmt.Println(line.Raw)
    }
}
```

### Check for an entry

```go
package main

import (
    "fmt"
    "github.com/ChrisWiegman/goodhosts"
)

func main() {
    hosts := goodhosts.NewHosts()

    if hosts.Has("127.0.0.1", "facebook.com") {
        fmt.Println("Entry exists!")
        return
    }

    fmt.Println("Entry doesn't exist!")
}
```

### Add an entry

```go
package main

import (
    "fmt"
    "github.com/ChrisWiegman/goodhosts"
)

func main() {
    hosts := goodhosts.NewHosts()

    // Note that nothing will be added to the hosts file until ``hosts.Flush`` is called.
    hosts.Add("127.0.0.1", "This is a line comment", "facebook.com", "twitter.com")

    if err := hosts.Flush(); err != nil {
        panic(err)
    }
}
```

### Remove an entry

```go
package main

import (
    "fmt"
    "github.com/ChrisWiegman/goodhosts"
)

func main() {
    hosts := goodhosts.NewHosts()

    // Same deal, yo: call hosts.Flush() to make permanent.
    hosts.Remove("127.0.0.1", "facebook.com", "twitter.com")

    if err := hosts.Flush(); err != nil {
        panic(err)
    }
}
```

### [More](API.md)

## Changelog

### 3.1.1 (2019-10-10)
* Fix existing tests

### 3.1 (2019-10-10)
* Allow sectioning of IP addresses with "section name" in api
* Various bugfixes

### 3.0.1 (2019-10-09)

* Refactored with go mod support
* Added ability to comment lines
* Only one host per line for easier management

### 2.1.0 (2015-06-08)

* Added Windows support.
* Added command-line docs.

### 2.0.0 (2015-05-04)

* Breaking API change.
* Add support for adding and removing multiple hosts.
* Added ``--all`` flag.
* Handle malformed IP addresses.

### 1.0.0 (2015-05-03)

- Initial release.

## License

[MIT](LICENSE)

<img src="http://static.messynessychic.com/wp-content/uploads/2013/08/rothschildparty2.jpg" width=400><br>
