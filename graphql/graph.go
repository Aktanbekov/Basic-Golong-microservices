package main

type Server struct {
	// accountClient 	*account.Clinet 
	// catalogClient	*catalog.Client
	// orderClinet		*order.Client
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
// 	accountClient, err := account.NewClient(accountUrl)
// 	if err != nil{
// 		return nil, err 
// 	}

// 	catalogClient, err := catalog.NewClinet(catalogUrl)
// 	if err != nil{
// 		accountClient.Close()
// 		return nil, err 
// 	}

// 	orderClient, err := order.NewClient(orderUrl)
// 	if err != nil{
// 		accountClient.Close()
// 		catalogClient.Close()
// 		return nil, err 
// 	}

	return &Server {
// 		accountClient,
// 		catalogClient,
// 		orderClient,
	}, nil
}


// func (s *Server) Mutation() MutationResolver{
// 	return &mutationResolver {
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver{
// 	return &queryResolver{
// 		server: s,

// 	}
// }


// func (s *Server) Account() AccountResolver{
// 	return &accountResolver{
// 		server: s,

// 	}
// }

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolver: s,
	})
}