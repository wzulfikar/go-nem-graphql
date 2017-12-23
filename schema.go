package nemgraphql

var GraphQLSchema = `
schema {
	query: Query
	mutation: Mutation
}

# The query type, represents all of the entry points into our object graph
type Query {
	# hello query is used to check if graphql is working correctly
	hello(name: String = "world!"): Hello!

	# retrieve transactions of given NEM address
	allTransactions(address: String!, hash: String = "", id: String = ""): TransactionsType!
}

type Mutation {
	newHello(name: String = "world!"): Hello!
}
`
