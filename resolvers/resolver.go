package nemgraphql

import (
	"log"

	nemparams "github.com/wzulfikar/go-nem-client/params"
	nemrequests "github.com/wzulfikar/go-nem-client/requests"
)

type Resolver struct {
	Client *nemrequests.Client
}

func (r *Resolver) Hello(args struct{ Name string }) *helloResolver {
	return &helloResolver{&hello{"hello", args.Name, "hash"}}
}

func (r *Resolver) AllTransactions(args nemparams.AllTransactions) []*transactionDataResolver {
	var l []*transactionDataResolver

	tx, err := r.Client.GetAllTransactions(args.Address, args.Hash, args.Id)
	if err != nil {
		log.Println(err)
		return l
	}

	for _, t := range tx.Data {
		l = append(l, &transactionDataResolver{&t})
	}

	return l
}
