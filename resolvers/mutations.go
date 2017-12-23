package nemgraphql

func (r *Resolver) NewHello(args struct{ Name string }) *helloResolver {
	defer Logger("Mutation newHello")()

	// Do some mutation here (store data to db, etc)..

	return &helloResolver{&hello{"new hello resolver id", args.Name, "hash"}}
}
