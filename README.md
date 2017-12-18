# Go NEM GraphQL
[WIP]

GraphQL for NEM network with GraphiQL page included.


### Installation

`go get -u github.com/wzulfikar/go-nem-graphql/cmd/nemgraphql`

> You need to install Go to have `go` command. See: https://golang.org/doc/install#install

### Running

Connect to NEM server at http://23.228.67.85:7890 and run the graphql on [http://localhost:3001](http://localhost:3001):

```
NEM_SERVER=http://23.228.67.85:7890 \
GRAPHQL_PORT=3001 \
nemgraphql
```

### Environment Variables

Here is the available environment variables and its corresponding default values:

```
NEM_SERVER=http://127.0.0.1:7890
GRAPHQL_PORT=8889
```
