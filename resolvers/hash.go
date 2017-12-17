package nemgraphql

import (
	nemtypings "github.com/wzulfikar/go-nem-client/typings"
)

type hashResolver struct {
	h *nemtypings.Hash
}

func (r *hashResolver) Data() *string {
	return &r.h.Data
}
