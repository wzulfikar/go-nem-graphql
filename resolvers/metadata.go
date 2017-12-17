package nemgraphql

import (
	nemtypings "github.com/wzulfikar/go-nem-client/typings"
)

type metaDataResolver struct {
	m *nemtypings.MetaData
}

func (r *metaDataResolver) Data() *string {
	return &r.m.Data
}
