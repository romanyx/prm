[![GoDoc](https://godoc.org/github.com/romanyx/prm?status.svg)](https://godoc.org/github.com/romanyx/prm)
[![Build Status](https://travis-ci.org/romanyx/prm.png)](https://travis-ci.org/romanyx/prm)
[![Coverage Status](https://coveralls.io/repos/github/romanyx/prm/badge.svg?branch=master)](https://coveralls.io/github/romanyx/prm?branch=master)

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