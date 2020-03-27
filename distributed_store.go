package main

import (
	"errors"
	"fmt"
)

type DistributedStore struct {
	stores []map[string]interface{}
	capacity int
}


func NewDistributedStore(count int) *DistributedStore {
	d := &DistributedStore{
		stores:   make([]map[string]interface{}, count),
		capacity: count,
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
