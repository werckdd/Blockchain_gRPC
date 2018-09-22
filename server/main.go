package main

import (
	"net"
	"log"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"qj/currency/cmd/Blockchain_gRPC/proto"
	"qj/currency/cmd/Blockchain_gRPC/server/blockchain"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("unable to listen on 8080 port: %v", err )
	}

	srv := grpc.NewServer()
	server := Server{
		Blockchain:blockchain.NewBlockchain(),
	}
	proto.RegisterBlockchainServer(srv, server)
	srv.Serve(listener)
}

type Server struct {
	Blockchain *blockchain.Blockchain

}

func (s Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)

	return &proto.AddBlockResponse{
		Hash:block.Hash,

	}, nil
}

func (s Server) GetBlockchain(ctx context.Context,in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash: b.Hash,
			Data: b.Data,
		})
	}
	return resp, nil
}

