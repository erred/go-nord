# go-nord [![GoDoc](https://godoc.org/github.com/seankhliao/go-nord-color?status.svg)](https://godoc.org/github.com/seankhliao/go-nord-color) [![Build Status](https://img.shields.io/travis/seankhliao/go-nord-color.svg?style=flat-square)](https://travis-ci.org/seankhliao/go-nord-color)

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
        fmt.Printf("These are the base colors: %v %v %v %v", nord.Nord0, nord.Nord1, nord.Nord2, nord.Nord3)
}
```
#### Output
```
These are the base colors: #2e3440 #3b4252 #434c5e #4c566a
```

## License

The MIT License (MIT) - see [`LICENSE.md`](LICENSE.md) for more details
