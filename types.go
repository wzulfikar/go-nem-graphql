package nemgraphql

var GraphQLTypes = `
type Hash {
	data: String
}

type Hello {
	id: String
	name: String
	hash: Hash
}

type TransactionMetaData {
	hash:  Hash
	height: String
	id:     String
}

type Transaction {
	deadline:      Int
	fee:           Int  
	mode:          Int  
	remoteAccount: String
	signature:     String
	signer:        String
	timeStamp:     Int
	type:          Int  
	version:       Int  
}

type TransactionData {
	meta: TransactionMetaData
	transaction: Transaction
}

type TransactionsType {
	totalCount: Int
	data: [TransactionData]
}
`
