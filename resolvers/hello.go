package nemgraphql

import (
	nemtypings "github.com/wzulfikar/go-nem-client/typings"
)

type hello struct {
	ID   string
	Name string
	Hash string
}

// Hello resolver is used for debugging purpose
type helloResolver struct {
	h *hello
}

func (r *helloResolver) ID() *string {
	return &r.h.ID
}

func (r *helloResolver) Name() *string {
	return &r.h.Name
}

func (r *helloResolver) Hash() *hashResolver {
	return &hashResolver{&nemtypings.Hash{r.h.Hash}}
}
