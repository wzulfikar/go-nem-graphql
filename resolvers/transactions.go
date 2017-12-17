package nemgraphql

type transactionsType struct {
	TotalCount int32
	Data       []*transactionDataResolver
}

type transactionsTypeResolver struct {
	t *transactionsType
}

func (r *transactionsTypeResolver) TotalCount() *int32 {
	return &r.t.TotalCount
}

func (r *transactionsTypeResolver) Data() *[]*transactionDataResolver {
	return &r.t.Data
}
