package cacheClient

type Cmd struct {
	Name  string
	Key   string
	Value string
	Error error
}

type Client interface {
	Run(cmd *Cmd)
}

func New(typ, server string) Client {
	if typ == "redis" {
		return newRedisClient(server)
	} else if typ == "http" {
		return newHTTPClient(server)
	} else if typ == "tcp" {
		return newTCPClient(server)
	}
	panic("Unknown client type: " + typ)
}
