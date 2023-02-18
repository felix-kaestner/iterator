# Iterator

<p align="center">
    <span>Go Iterator Implementation with support for Generics (requires Go v1.18+).</span>
    <br><br>
    <a href="https://github.com/felix-kaestner/iterator/issues">
        <img alt="Issues" src="https://img.shields.io/github/issues/felix-kaestner/iterator?color=%2329b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/iterator/stargazers">
        <img alt="Stars" src="https://img.shields.io/github/stars/felix-kaestner/iterator?color=%2329b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/iterator/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/github/license/felix-kaestner/iterator?color=%2329b6f6&style=flat-square"/>
    </a>
    <a href="https://pkg.go.dev/github.com/felix-kaestner/iterator">
        <img alt="Stars" src="https://img.shields.io/badge/go-documentation-blue?color=%2329b6f6&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/felix-kaestner/iterator">
        <img alt="Issues" src="https://goreportcard.com/badge/github.com/felix-kaestner/iterator?style=flat-square">
    </a>
    <a href="https://codecov.io/gh/felix-kaestner/iterator">
        <img src="https://img.shields.io/codecov/c/github/felix-kaestner/iterator?style=flat-square&token=KK7ZG7A90X"/>
    </a>
    <a href="https://twitter.com/kaestner_felix">
        <img alt="Twitter" src="https://img.shields.io/badge/twitter-@kaestner_felix-29b6f6?style=flat-square">
    </a>
</p>

## Quickstart

```go
package main

import (
	"fmt"

	"github.com/felix-kaestner/iterator"
)

func main() {
	s := []int{1, 2, 3}
	i := iterator.FromSlice(s)

	for i.HasNext() {
		item, _ := i.Next()
		fmt.Println(fmt.Sprintf("item: %d", *item))
	}
}
```

Output:

```sh
item: 1
item: 2
item: 3
```

##  Installation

Install with the `go get` command:

```
$ go get -u github.com/felix-kaestner/iterator
```

## Contribute

All contributions in any form are welcome! 🙌🏻  
Just use the [Issue](.github/ISSUE_TEMPLATE) and [Pull Request](.github/PULL_REQUEST_TEMPLATE) templates and I'll be happy to review your suggestions. 👍

## Cheers ✌🏻
