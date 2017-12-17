package nemgraphql

import (
	"strconv"

	nemtypings "github.com/wzulfikar/go-nem-client/typings"
)

type transactionMetaDataResolver struct {
	t *nemtypings.TransactionMetaData
}

func (r *transactionMetaDataResolver) Hash() *hashResolver {
	return &hashResolver{&r.t.Hash}
}

func (r *transactionMetaDataResolver) Height() *string {
	i := strconv.Itoa(int(r.t.Height))
	return &i
}

func (r *transactionMetaDataResolver) ID() *string {
	i := strconv.Itoa(int(r.t.ID))
	return &i
}
