package server

import (
	"fmt"
	"net/http"
	"server/internal/config"
	"server/internal/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

// Server struct to manage application server
type Server struct{}

// NewServer initializes a new server instance
func GetServer() *Server {
	return &Server{}
}

type ServerArgs struct {
	DB *gorm.DB
}

// Start runs the HTTP server
func (s *Server) Start(args ServerArgs) error {
	// GraphQL handler
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: args.DB,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Define HTTP routes
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("🚀 Server running at http://localhost:" + config.Cfg.ServerPort)
	return http.ListenAndServe(":"+config.Cfg.ServerPort, nil)
}
