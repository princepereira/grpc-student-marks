package main

import (
	"context"
	"flag"
	"fmt"
	pbproto "grpc-student-marks/pkg/proto-pb"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

// Ref: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/client/client.go

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

// Invoking rpc call towards the server
func sendGetAllMarksReq(ctx context.Context, client pbproto.StudentServiceClient) {
	fmt.Println("\n>>>>>>>>>>>>>>.sendGetAllMarksReq")
	req := new(pbproto.MarkReq)
	subjectInfo := new(pbproto.SubjectInfo)
	subjectInfo.SlNo = 123
	subjectInfo.Name = "Prince Pereira"
	req.SubjectInfo = subjectInfo
	btes, _ := proto.Marshal(req)
	fmt.Printf("\nMarshalled Data : %v\n", btes)
	fmt.Printf("Get-all-marks-for-Prince : %v\n", req)
	resp, err := client.GetAllMarks(ctx, req)

	if err != nil {
		log.Fatalf("could not get all marks: %v", err)
	}

	log.Printf("\nAll-Marks-for-prince: %s\n", resp)
}

// Invoking rpc call towards the server
func sendGetMarkReq(ctx context.Context, client pbproto.StudentServiceClient) {
	fmt.Println("\n>>>>>>>>>>>>>>.sendGetMarkReq")
	req := new(pbproto.MarkReq)
	subjectInfo := new(pbproto.SubjectInfo)
	subjectInfo.SlNo = 123
	subjectInfo.Name = "Prince Pereira"
	subjectInfo.Subject = pbproto.Subject_PHYSICS
	req.SubjectInfo = subjectInfo
	btes, _ := proto.Marshal(req)
	fmt.Printf("Marshalled Data : %v\n", btes)
	fmt.Printf("Get-mark-for-Prince : %v\n", req)
	resp, err := client.GetMark(ctx, req)

	if err != nil {
		log.Fatalf("could not get marks: %v\n", err)
	}

	log.Printf("Mark-for-prince: %s\n", resp)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	// Establishing the server connection
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Constructing a client object
	client := pbproto.NewStudentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if ctx == nil || client == nil {
		fmt.Println("Conext or client is nil")
	}

	sendGetMarkReq(ctx, client)
	sendGetAllMarksReq(ctx, client)
}
