package main

import (
	"context"
	"flag"
	"fmt"
	pbproto "grpc-student-marks/pkg/proto-pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Ref: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go

var (
	port = flag.Int("port", 10000, "The server port")
)

// Server struct for invoking server calls
type Server struct {
}

// GetMark refers to the method getting called for RPC call invoked for Get Single Subject Mark.
func (server *Server) GetMark(ctx context.Context, req *pbproto.MarkReq) (*pbproto.MarkResp, error) {
	fmt.Printf("Request received for GetMark() request. Req : %v \n", req)
	resp := new(pbproto.MarkResp)

	subjectInfo := new(pbproto.SubjectInfo)
	subjectInfo.SlNo = req.SubjectInfo.SlNo
	subjectInfo.Name = req.SubjectInfo.Name
	subjectInfo.Subject = req.SubjectInfo.Subject
	resp.SubjectInfo = subjectInfo

	if req.SubjectInfo.Subject == pbproto.Subject_PHYSICS {
		resp.Mark = 98
	} else {
		resp.Mark = 97
	}
	fmt.Printf("Response send for GetMark() request. Req : %v , Resp : %v\n", req, resp)
	return resp, nil
}

// GetAllMarks refers to the method getting called for RPC call invoked for Get all Subject Marks.
func (server *Server) GetAllMarks(ctx context.Context, req *pbproto.MarkReq) (*pbproto.AllMarksResp, error) {
	fmt.Printf("Request received for GetAllMarks() request. Req : %v \n", req)
	resp := new(pbproto.AllMarksResp)
	resp.SlNo = req.SubjectInfo.SlNo
	resp.Name = req.SubjectInfo.Name

	var markInfo []*pbproto.AllMarksResp_MarkInfo

	phyMark := new(pbproto.AllMarksResp_MarkInfo)
	phyMark.SubjectName = 0
	phyMark.Mark = 98
	markInfo = append(markInfo, phyMark)

	mathsMark := new(pbproto.AllMarksResp_MarkInfo)
	mathsMark.SubjectName = 2
	mathsMark.Mark = 97
	markInfo = append(markInfo, mathsMark)

	resp.Marks = markInfo

	fmt.Printf("Response send for GetAllMarks() request. Req : %v , Resp : %v\n", req, resp)
	return resp, nil
}

func main() {
	flag.Parse()
	// Listening to the port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Constructing a grpc server object.
	grpcServer := grpc.NewServer()

	pbproto.RegisterStudentServiceServer(grpcServer, &Server{})
	// Starting the server
	grpcServer.Serve(lis)
}
