# Strings-go

---

Tools-go is a tool library for Golang.

## License

Apache License, Version 2.0

## Implements

### TernaryOperator

* TernaryOperator
* TernaryOperatorNest

### ReflectSupport

* UnpackType
* UnpackValue

## Usage

### Installation

```shell
go get -u github.com/changedenczd/go-lib/tools
```

### Demo

#### TernaryOperator

```go
package main

import (
	"fmt"
	"time"
)

import (
	"github.com/changedenczd/go-lib/tools"
)

func main() {
	condition := (time.Now().UnixNano()/1e6)%2 == 1
	trueResult := "true"
	falseResult := "false"
	result := tools.TernaryOperator(condition, trueResult, falseResult)
	fmt.Println(result)
}
```
