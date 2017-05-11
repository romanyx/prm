[![GoDoc](https://godoc.org/github.com/romanyx/prm?status.svg)](https://godoc.org/github.com/romanyx/prm)
[![Build Status](https://travis-ci.org/romanyx/prm.png)](https://travis-ci.org/romanyx/prm)

# Overview

`prm` is a small golang library that replaces special characters in a string so that it may be used as part of a ‘pretty’ URL. Slug generator. It's very fast, simple and configurable.

# Usage

```go
package main

import (
	"fmt"

	"github.com/romanyx/prm"
)

func main() {
	prm.Parameterize("Computer world", '-') // "computer-world"
}
```