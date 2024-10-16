package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)


func main() {
	mc := memcache.New("localhost:11211")
	err := mc.Set(&memcache.Item{
		Key: "username",
		Value: []byte("imanimen"),
	})

	if err != nil {
		fmt.Println("Set error: ", err)
	}

	item, err := mc.Get("username")
	if err != nil {
		fmt.Println("Get error: ", err)
	}
	fmt.Println("memcache: Value:", string(item.Value))
}