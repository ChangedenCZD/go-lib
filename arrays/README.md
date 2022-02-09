# Strings-go

---

Arrays-go is a array extension library for Golang.

## License

Apache License, Version 2.0

## Implements

### Copy

* Copy
* CopyOf
* CopyOfRange
* CopyFrom

## Usage

### Installation

```shell
go get -u github.com/changedenczd/go-lib/arrays
```

### Demo

#### Copy

```go
package main

import (
	"fmt"
)

import (
	arrays "github.com/changedenczd/go-lib/arrays"
)

func main() {
	src := []interface{}{"1", 2, 3.14, false, func() {}}
	dest := make([]interface{}, len(src))
	arrays.CopyFrom(src, 1, dest, 2, 3)
	j := 1
	for i := 2; i < len(dest); i++ {
		fmt.Println(src[j], dest[i])
		j++
	}
}
```
