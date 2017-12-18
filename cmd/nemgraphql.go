package main

import (
	nemclient "local-dev/go-nem-client"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	nemgraphql "github.com/wzulfikar/go-nem-graphql"
	resolvers "github.com/wzulfikar/go-nem-graphql/resolvers"
)

var schema *graphql.Schema

var port = os.Getenv("GRAPHQL_PORT")
var nemServer = os.Getenv("NEM_SERVER")

// set default values
func init() {
	if port == "" {
		port = "8889"
	}
	if nemServer == "" {
		nemServer = "http://127.0.0.1:7890"
	}
}

func main() {
	c, err := nemclient.NewClient(nemServer)
	if err != nil {
		log.Fatal(err)
	}

	resolver := &resolvers.Resolver{
		Client: c,
	}

	graphqlSchema := strings.Join([]string{
		nemgraphql.GraphQLSchema,
		nemgraphql.GraphQLTypes,
	}, "")
	schema = graphql.MustParseSchema(graphqlSchema, resolver)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(nemgraphql.GraphiQLPage))
	}))
	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Println("→ Connecting to NEM server at", nemServer)
	log.Println("✔ NEM GraphQL started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
