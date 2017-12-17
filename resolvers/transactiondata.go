package nemgraphql

import (
	nemtypings "github.com/wzulfikar/go-nem-client/typings"
)

type transactionDataResolver struct {
	t nemtypings.TransactionData
}

func (r *transactionDataResolver) Meta() *transactionMetaDataResolver {
	return &transactionMetaDataResolver{&r.t.Meta}
}

func (r *transactionDataResolver) Transaction() *transactionResolver {
	return &transactionResolver{&r.t.Transaction}
}
