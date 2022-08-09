# Common array utilities for golang arrays

* Every - Checks if every element in an array pass a test
* Filter - Creates a new array with every element in an array that pass a test
* Find - Returns the value of the first element in an array that pass a test
* Map - Creates a new array with the result of calling a function for each array element
* Some - Checks if any of the elements in an array pass a test
* Reduce - Reduce the values of an array to a single value (going left-to-right)

## Installation

```sh
go get github.com/judedaryl/go-arrayutils
```

## Usage

```go
package main

import "github.com/judedaryl/go-arrayutils"

func Even(val int) bool {
	return val%2 != 0
}

func main() {
	myArr := []int{1, 2, 3, 4}
	myArr = arrayutils.Filter(myArr, Even)
	content, _ := json.Marshal(myArr)
	log.Println(string(content))
}
```
