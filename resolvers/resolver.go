package nemgraphql

import (
	"fmt"
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

func (r *Resolver) AllTransactions(args nemparams.AllTransactions) *transactionsTypeResolver {
	tx, err := r.Client.GetAllTransactions(args.Address, args.Hash, args.Id)
	if err != nil {
		log.Println(err)
		return &transactionsTypeResolver{}
	}

	var l []*transactionDataResolver
	for _, t := range tx.Data {
		l = append(l, &transactionDataResolver{t})
	}

	fmt.Println(len(tx.Data), "tx found")

	return &transactionsTypeResolver{
		&transactionsType{
			TotalCount: int32(len(tx.Data)),
			Data:       l,
		},
	}
}
