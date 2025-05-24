# EcommerceV2

## What is GraphQL?
GraphQL is a query language for APIs, and a server side runtime for executing queries using a type system you define for your data.

This means it provides a way to request and modify data from multiple data sources (for instance, a REST API, SQLite database, and services) in a single query.

query that specifies exactly what data you need.

There are two main components of GraphQL:
1.- GraphQL Server
2.- GraphQL Client

The GraphQL server is responsible for implementing the GraphQL API on the server side. Examples of GrapQL servers are:
1.- Express
2.- Apollo server

The GraphQL client allows apps to interact with the GraphQL server. It enables fetching and updating data in a declarative manner. Examples of GraphQL clients are:
1.- Relay
2.- Apollo client

GraphQL API is built on top of a GraphQL server. The server is the central point for receiving GraphQL queries. Once the query is received, it will be matched against a defined schema. The server retrieves the data by interacting with databsaes, microservices or other data sources.

The GraphQL server consists of:
1.- Schema Definition Language (SDL): This represents the structure or shape of your data set. The SDL is used to define types, fields, and relationships between types.

2.- Resolvers: They are functions that specify how to process specific GraphQL operations and contain the logic for fetching the requested data. These functions can be Queries or Mutations. Resolvers can fetch data from databases, external APIs, or any other data source.

3.- Data Sources: Any available data source. For instance, MySQL, MongoDB, etc.

    Server  -->     Schema
                    Resolver
                    Data Sources

The GraphQL client sends queries and mutations to the server, and receive the response.
1.- Queries: These are similar to REST APIs enpoints where client hits the request an get some piece of information in response (by the HTTP GET Method), similarly in GraphQL. Queries are used by clients to request specfic data from GraphQL server. They resemble the shape of the data that the client expects to receive. Queries are executed against the GraphQL schema, and the server responds with the requested data.

**But unlike REST API, in Graphql query client needs to tell which information it wants to be fetched by server as the below mentioned query. That is to say, the client needs to specify through query what data the server have to response.**

2.- Mutations: Mutations are used to modify data in the server. Similar to The RESTful APIs utilize HTTP verbs to perform CRUD (Create Read Update Delete) operations. While the read operation (GET) is made possible through the use of queries in GraphQL, the rest three operations (POST, PUT, DELETE) are carried out through mutations. They allow clients to create, update, or delete data. Like queries, mutations are executed against the GraphQL schema.

Go to:
    https://medium.com/@naveengupta262001/graphql-core-components-in-detail-04e93487b0ba

# Dependencies
go mod init github.com/luisrojas17/ecommerceV2

go get github.com/99designs/gqlgen

cd  graphql
go run github.com/99designs/gqlgen generate

Delete next files:

    generated.go
    models_generated.go

Also, you will have to get next modules to config EntryPoint in main.go script.

    go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.70
    go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.70
    go get github.com/99designs/gqlgen/handler@v0.17.70
    go get github.com/kelseyhightower/envconfig
    go get github.com/lib/pq
    go get github.com/segmentio/ksuid
    go get github.com/tinrab/retry
    go get -u google.golang.org/grpc
    go get gopkg.in/olivere/elastic.v5

Or execute next command to clean and updates "go.mod" and "go.sum" by ensuring they reflect the actual dependencies required by the code.

    go mod tidy

## Install Proto Plug-Ins

wget https://github.com/protocolbuffers/protobuf/releases/download/v{version}/protoc-{version}-linux-x86_64.zip

1. wget https://github.com/protocolbuffers/protobuf/releases/download/v31.0/protoc-31.0-linux-x86_64.zip
2. unzip protoc-31.0-linux-x86_64.zip -d protoc
3. sudo mv protoc/bin/protoc /usr/local/bin/

sudo apt install protobuf-compiler
4. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
5. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
6. echo $PATH
7. export PATH="$PATH:$(go env GOPATH)/bin"
8. source ~/.bashrc
9. Create the "pb" folder inside "accounts" folder
10. Add next line to account.proto file: option go_package = "github.com/luisrojas17/ecommerceV2/accounts/pb";
11. Run this command: protoc --go_out=./pb --go-grpc_out=./pb accounts.proto

# Run

cd  graphql
go run main.go

# Microservice Flow and Description

GrapQL
    Queries     Products
                Accounts
                Orders

    Mutations   Create Account
                Create Product
                Create Order


Mutation/Query -> Client -> Server -> Service -> Respository -> Database