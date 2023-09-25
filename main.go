package main

// func main() {
// 	user := graphql.NewObject(graphql.ObjectConfig{
// 		Name: "User",
// 		Fields: graphql.Fields{
// 			"name": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"description": &graphql.Field{
// 				Type: graphql.String,
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					user, ok := p.Source.(User)
// 					if ok {
// 						fmt.Println("user:", user)
// 					} else {
// 						fmt.Println(p.Source)
// 					}
// 					return "descript", nil
// 				},
// 			},
// 			"id": &graphql.Field{
// 				Type: graphql.Int,
// 			},
// 			"otherNames": &graphql.Field{
// 				Type: graphql.NewList(graphql.String),
// 			},
// 			"imageUrl": &graphql.Field{
// 				Type: graphql.String,
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					return "url", nil
// 				},
// 			},
// 		},
// 	})
// 	user.AddFieldConfig("mother", &graphql.Field{
// 		Type: user,
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 			fmt.Println(p.Source)
// 			return User{Name: "mother"}, nil
// 		},
// 	})
// 	fields := graphql.Fields{
// 		"user": &graphql.Field{
// 			Type: user,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return User{Name: "child"}, nil
// 			},
// 		},
// 		"hello": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return "world", nil
// 			},
// 		},
// 	}
// 	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
// 	var RootSubscription = graphql.NewObject(graphql.ObjectConfig{
// 		Name: "RootSubscription",
// 		Fields: graphql.Fields{
// 			"feed": &graphql.Field{
// 				Type: user,
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					return p.Source, nil
// 				},
// 				Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
// 					c := make(chan interface{})
// 					go func() {
// 						var i int
// 						for {
// 							i++

// 							feed := User{Name: "sub"}

// 							select {
// 							case <-p.Context.Done():
// 								log.Println("[RootSubscription] [Subscribe] subscription canceled")
// 								close(c)
// 								return
// 							default:
// 								c <- feed
// 							}
// 							time.Sleep(250 * time.Millisecond)
// 							if i == 21 {
// 								close(c)
// 								return
// 							}
// 						}
// 					}()

// 					return c, nil
// 				},
// 			},
// 		},
// 	})
// 	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Subscription: RootSubscription}
// 	schema, err := graphql.NewSchema(schemaConfig)
// 	if err != nil {
// 		log.Fatalf("failed to create new schema, error: %v", err)
// 	}
// 	h := handler.New(&handler.Config{
// 		Schema:   &schema,
// 		Pretty:   true,
// 		GraphiQL: false,
// 	})
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte("{\"hello\": \"world\"}"))
// 	})
// 	mux.Handle("/graphql", h)
// 	handler := cors.Default().Handler(mux)

// 	http.ListenAndServe(":8080", handler)

// }
