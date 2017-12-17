package nemgraphql

import (
	nemtypings "github.com/wzulfikar/go-nem-client/typings"
)

type transactionResolver struct {
	t *nemtypings.Transaction
}

func (r *transactionResolver) Deadline() *int32 {
	i := int32(r.t.Deadline)
	return &i
}

func (r *transactionResolver) Fee() *int32 {
	i := int32(r.t.Fee)
	return &i
}

func (r *transactionResolver) Mode() *int32 {
	i := int32(r.t.Mode)
	return &i
}

func (r *transactionResolver) RemoteAccount() *string {
	return &r.t.RemoteAccount
}
func (r *transactionResolver) Signature() *string {
	return &r.t.Signature
}

func (r *transactionResolver) Signer() *string {
	return &r.t.Signer
}

func (r *transactionResolver) TimeStamp() *int32 {
	i := int32(r.t.TimeStamp)
	return &i
}

func (r *transactionResolver) Type() *int32 {
	i := int32(r.t.Type)
	return &i
}

func (r *transactionResolver) Version() *int32 {
	i := int32(r.t.Version)
	return &i
}
