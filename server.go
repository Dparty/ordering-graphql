package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Dparty/common/utils"
	"github.com/Dparty/model"
	coreModel "github.com/Dparty/model/core"
	"github.com/Dparty/ordering-graphql/graph"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			claims, _ := utils.VerifyJwt(authorization)
			account, _ := coreModel.FindAccount(utils.StringToUint(claims["id"].(string)))
			ctx := context.WithValue(r.Context(), userCtxKey, account)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	var err error
	viper.SetConfigName(".env.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	db, err := model.NewConnection(user, password, host, port, database)
	if err != nil {
		panic(err)
	}
	graph.Init(db)
	port = os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{})
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
