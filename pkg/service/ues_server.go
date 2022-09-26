package service

import (
	"fmt"
	"github.com/wangxn2015/go_simple_grpc_ues/api/track_msg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UesServer struct {
	track_msg.UnimplementedUEsServer
	ueStore UeStore
}

func NewUesServer(ueStore UeStore) *UesServer {
	return &UesServer{
		ueStore: ueStore,
	}
}

func (server *UesServer) GetUEs(req *track_msg.UERequest, stream track_msg.UEs_GetUEsServer) error {

	fmt.Println("receive GetUEs request: %v ", req)
	err := server.ueStore.Search(
		stream.Context(),
		func(ue *track_msg.UEInfo) error {
			//res := &track_msg.UEInfo{Laptop: ue}
			res := ue
			err := stream.Send(res)
			if err != nil {
				return err
			}

			fmt.Printf("sent ue with id: %d\n", ue.GetImsi())
			return nil
		},
	)

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
	//return status.Errorf(codes.Unimplemented, "method GetUEs not implemented")
}
