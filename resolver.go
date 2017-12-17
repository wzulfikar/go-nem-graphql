package nemgraphql

import nemparams "github.com/wzulfikar/go-nem-client/params"

type Resolver struct{}

func (r *Resolver) AllTransactions(args nemparams.AllTransactions) *string {
	temp := "all tx" + args.Address + args.Hash + args.Id
	return &temp
}
