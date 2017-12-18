package nemgraphql

var GraphQLSchema = `
schema {
	query: Query
	# mutation: Mutation
}

# The query type, represents all of the entry points into our object graph
type Query {
	hello(name: String = "world!"): Hello!
	allTransactions(address: String!, hash: String = "", id: String = ""): TransactionsType!
}
`
