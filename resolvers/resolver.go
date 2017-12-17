package nemgraphql

import (
	nemrequests "github.com/wzulfikar/go-nem-client/requests"
)

type Resolver struct {
	Client *nemrequests.Client
}

func (r *Resolver) Hello(args struct{ Name string }) *helloResolver {
	return &helloResolver{&hello{"my id", args.Name, "hash"}}
}
