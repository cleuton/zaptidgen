// Copyright 2023 Cleuton Sampaio.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/cleuton/zaptidgen/gen"

	"github.com/sony/sonyflake"
	"google.golang.org/grpc"
)

// ***** Global vars *****
const VERSION = "0.1.0"

var logger log.Logger = *log.Default()
var st sonyflake.Settings
var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

// ***** Protobuf server *****
type server struct {
	pb.UnimplementedIdGenServer
}

func (s *server) Gen(ctx context.Context, input *pb.IdRequest) (*pb.IdResponse, error) {
	errorCode := false
	sequence := uint64(0)
	var returnError error = nil
	if flake == nil {
		logger.Panicf("FATAL ERROR: Couldn't generate sonyflake.NewSonyflake.\n")
	}
	sequence, err := flake.NextID()
	if err != nil {
		errorCode = true
		returnError = errors.New("NextID failed")
		log.Printf("ERROR: flake.NextID() failed with %s\n", err)
	}
	return &pb.IdResponse{Error: errorCode, Id: sequence}, returnError
}

// This is zapt.chat ID generator server.
// A gRPC server that exports unique sequence of uint64 numbers.
// It uses the last 16 bits of machine's IP address as "machine number".
func main() {
	logger.Printf("INFO: ZAPTIDGEN - V %s\n - starting... (port: 8888)", VERSION)

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		logger.Panicf("FATAL ERROR: failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterIdGenServer(s, &server{})
	logger.Printf("INFO: server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Panicf("FATAL ERROR: failed to serve: %v\n", err)
	}
}
