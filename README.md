# go-nord [![GoDoc][1]][2] [![Build Status][3]][4] [![Go Report Card][5]][6] [![License: MIT][7]][8]

[1]: https://godoc.org/github.com/seankhliao/go-nord?status.svg
[2]: https://godoc.org/github.com/seankhliao/go-nord
[3]: https://img.shields.io/travis/seankhliao/go-nord.svg?style=flat-square
[4]: https://travis-ci.org/seankhliao/go-nord
[5]: https://goreportcard.com/badge/github.com/seankhliao/go-nord?style=flat-square
[6]: https://goreportcard.com/report/github.com/seankhliao/go-nord
[7]: https://img.shields.io/badge/License-MIT-blue.svg?longCache=true&style=flat-square
[8]: LICENSE

nord provides the [nord](https://github.com/arcticicestudio/nord) colors for use in Go

## Install

```bash
go get github.com/seankhliao/go-nord
```

## Example

```go
package main

import (
        "github.com/seankhliao/go-nord"
        "fmt"
)

func main() {
        fmt.Printf("These are the base colors: %v %v %v %v", 
        nord.Nord0, nord.Nord1, nord.Nord2, nord.Nord3)
}
```
#### Output
```
These are the base colors: #2e3440 #3b4252 #434c5e #4c566a
```
