package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	nemgraphql "github.com/wzulfikar/go-nem-graphql"
)

var schema *graphql.Schema

var port = os.Getenv("GRAPHQL_PORT")
var schemaFile = os.Getenv("GRAPHQL_SCHEMA")
var nemServer = os.Getenv("NEM_SERVER")

// set default values
func init() {
	if schemaFile == "" {
		schemaFile = "../schema.graphql"
	}
	if port == "" {
		port = "8080"
	}
	if nemServer == "" {
		nemServer = "http://127.0.0.1:7890"
	}
}

func main() {
	schema = graphql.MustParseSchema(string(file(schemaFile)), &nemgraphql.Resolver{})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(nemgraphql.GraphiQLPage))
	}))
	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func file(filename string) []byte {
	path, _ := filepath.Abs(filename)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
