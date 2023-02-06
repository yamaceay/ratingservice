// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	// "github.com/golang/protobuf/proto"
	pb "github.com/yamaceay/ratingservice/src/ratingservice/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"fmt"
	"proto"
	"reflect"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	addr := run(port)
	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure())

	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewRatingServiceClient(conn)

	got, err := client.GetRatings(ctx, &pb.GetRatingsRequest{ProductId: "OLJCESPC7Z"})
	if err != nil {
		t.Fatal(err)
	}
	parsedRatings, _ := parseRatings()
	want := parsedRatings["OLJCESPC7Z"]
	fmt.Printf("Wanted parsed from ratings: %v", want)
	fmt.Printf("Type of want %v", reflect.TypeOf(want))
	fmt.Printf("Type of got %v", reflect.TypeOf(got))

	if proto.Equal(got, want[0]){
		fmt.Print("This was correct")
	}

	if want, _ := parseRatings(); 1 == 1{
		t.Errorf("got %v, want %v", got, want)
	}
	got1, err1 := client.GetRatings(ctx, &pb.GetRatingsRequest{ProductId: "66VCHSJNUP"})
	if err1 != nil {
		t.Fatal(err1)
	}
	if want1, _ := parseRatings(); 1 == 1{
		t.Errorf("got %v, want %v", got1, want1)
	}
	_, err = client.GetRatings(ctx, &pb.GetRatingsRequest{ProductId: "N/A"})
	if got, want := status.Code(err), codes.OK; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
