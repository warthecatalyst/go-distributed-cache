package main

import (
	"fmt"
	"go-distributed-cache/http_in_memory_cache/cache"
)

func main() {
	c := cache.New("inmemory")
	buf := []byte{1, 2, 3}
	fmt.Print(c.Set("123", buf))
}
