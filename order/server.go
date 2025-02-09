package order

import (
	"fmt"
	"log"

	"github.com/Aktanbekov/Basic-Golong-microservices/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type grpcServer struct {
	service		Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		s,
		accountClient,
		catalogClient,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostOrder(ctx context.Context, r *pb.PostOrderRequest)(*pb.PostOrderResponse, error) {
	_, err := s.accountClinet.GetAccount(ctx, r.AccountId)
	if err != nil {
		log.Println("Error getting account:", err)
		return nil, errors.New("account not found")
	}

	productIDs := []string{}
	orderedProducts, err := s.catalog.GetProducts(ctx, 0, 0, productIDs, "")
	if err != nil {
		log.Println("Error getting products: ", err)
		return nil, errors.New("products not found")
	}
	products := []OrderedProduct{}
	for _, p := range orderedProducts{
		product := OrderedProducts{
			ID:		 	p.ID,
			Quantity: 	0,
			Price: 		p.Prince,
			Name: 		p.name,
			Description:p.Description,
		}
		for _, rp := range r.Products{
			if rp.ProductId == p.ID{
				product.Quantity = rp.Quantity
				break
			}
		}
		if product.Quantity != 0 {
			products = append(products, product)
		}
	} 
	order, err := s.service.PostOrder(ctx, r.AccountId, products)
	if err != nil {
		log.Println("Error posting order:", err)
		return nil, errors.New("could not post order")
	}

	orderProto := &pb.Order{
		Id: order.ID,
		AccountId: order.AccountID,
		TotalPrice: order.TotalPrice,
		Products: []*pb.Order_OrderProduct{},
	}
	orderProto.CreatedAt, _ = order.CreatedAt
}