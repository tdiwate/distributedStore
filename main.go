package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type DistributedStore struct {
	stores []map[string]interface{}
	capacity int
}


func NewDistributedStore(count int) *DistributedStore {
	d := &DistributedStore{
		stores:   make([]map[string]interface{}, 10),
		capacity: 10,
	}
	for i := range d.stores {
		d.stores[i] = make(map[string]interface{})
	}
	return d
}

func (d *DistributedStore) Write(key string, value interface{})  {
	loc := hashFunc(key, d.capacity)
	d.stores[loc][key] = value
}

func (d *DistributedStore) Read(key string) (interface{}, error){
	loc := hashFunc(key, d.capacity)
	_, ok := d.stores[loc][key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return d.stores[loc][key], nil
}

func (d *DistributedStore) Dump() {
	for i := range d.stores {
		fmt.Println("Store: ", i)
		for key, val := range d.stores[i] {
			fmt.Println(key, val)
		}
	}
}

func main() {
	dStore := NewDistributedStore(10)
	dStore.Write("apple", 1)
	dStore.Write("banana", 2)
	dStore.Dump()
	v, err := dStore.Read("apple")
	fmt.Println(v, err)
	v, err = dStore.Read("hello")
	fmt.Println(v, err)
}

func hashFunc(s string, mod int) int {
	loc :=0
	for _, i := range s {
		loc += int(i)
	}
	return loc%mod
}
