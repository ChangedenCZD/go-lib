# Collection-go

---

Collection-go is a collection extension library for Golang.

## License

Apache License, Version 2.0

## Implements

### Vector

* Get
* Set
* Add
* Remove
* Size
* IsEmpty
* AddElement
* SetElementAt
* InsertElementAt
* RemoveElement
* RemoveElementAt
* RemoveAllElements
* ElementAt
* FirstElement
* LastElement
* Clone
* Contains
* IndexOf
* LastIndexOf
* Elements

### Stack

extends Vector

* Push
* Pop
* Peek
* Search
* Empty

## Usage
### Installation
```shell
go get -u github.com/changedenczd/go-lib/collections
```

### Demo
#### Vector

```go
package main

import (
	"fmt"
)

import (
	collection "github.com/changedenczd/go-lib/collections"
)

func main() {
	vector := collection.NewVector()
	vector.Add(100)
	vector.AddElement(200)
	vector.AddElement(300)

	vector.Remove(1)
	item := vector.LastElement()
	fmt.Println(item)
	item = vector.FirstElement()
	fmt.Println(item)

	vector.RemoveElementAt(0)
	vector.RemoveAllElements()
}
```

#### Stack

```go
package main

import (
	"fmt"
)

import (
	collection "github.com/changedenczd/go-lib/collections"
)

func main() {
	stack := collection.NewStack()
	stack.Push(1)
	stack.Push(2)
	peek := stack.Peek()
	fmt.Println(peek)
	pop := stack.Pop()
	fmt.Println(pop)
}
```
