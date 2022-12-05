package main

import (
	"go-distributed-cache/http_in_memory_cache/cache"
	"go-distributed-cache/http_in_memory_cache/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
