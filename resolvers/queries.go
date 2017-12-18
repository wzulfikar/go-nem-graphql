package nemgraphql

import (
	"errors"
	"fmt"
	"log"
	"strings"

	nemparams "github.com/wzulfikar/go-nem-client/params"
)

func (r *Resolver) Hello(args struct{ Name string }) *helloResolver {
	return &helloResolver{&hello{"hello", args.Name, "hash"}}
}

func (r *Resolver) AllTransactions(args nemparams.AllTransactions) (*transactionsTypeResolver, error) {
	defer Logger("Query AllTransactions")()

	tx, err := r.Client.GetAllTransactions(args.Address, args.Hash, args.Id)
	if err != nil {
		log.Println("error", err)
		if strings.Contains(err.Error(), "connection refused") {
			return nil, errors.New("NEM server error: can't connect to " + r.Client.Endpoint)
		}
		return nil, errors.New("Something went wrong :(")
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
	}, nil
}
