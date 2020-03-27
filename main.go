package main

import (
	"fmt"
)

func main() {
	dStore := NewDistributedStore(10)
	dStore.Write("apple", "1.00")
	dStore.Write("banana", 200)
	dStore.Write("orange", 1.34)
	dStore.Write("pineapple", 20)
	dStore.Write("banana", 5)
	dStore.Dump()
	v, err := dStore.Read("apple")
	fmt.Println(v, err)
	v, err = dStore.Read("guava")
	fmt.Println(v, err)
}
