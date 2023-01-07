package tcp_in_memory_cache

import (
	"go-distributed-cache/http_in_memory_cache/cache"
	"go-distributed-cache/http_in_memory_cache/http"
	"go-distributed-cache/tcp_in_memory_cache/tcp"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
