package tests

import (
	"context"
	"fmt"
	"testing"

	"go-micro-frame/proto"
)

func init() {
	InitClient()
}

func Test_GetUserList(t *testing.T) {
	rsp, err := GrpcClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		if err != nil {
			panic(err)
		}
	}

	defer ClientConn.Close()
}

func Test_GetUserById(t *testing.T) {
	rsp, err := GrpcClient.GetUserById(context.Background(), &proto.IdRequest{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)

	defer ClientConn.Close()
}
