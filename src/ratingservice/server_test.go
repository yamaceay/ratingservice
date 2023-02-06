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
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	"testing"
	// "fmt"
	"github.com/golang/protobuf/proto"
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

	for i, _ := range got.Ratings {
		if !proto.Equal(got.Ratings[i], want[i]) {
			t.Errorf("\n%v\n%v\n", got.Ratings[i], want[i])
		}
	}

	got1, err1 := client.GetRatings(ctx, &pb.GetRatingsRequest{ProductId: "66VCHSJNUP"})
	parsedRatings1, _ := parseRatings()
	want1 := parsedRatings1["66VCHSJNUP"]
	if err1 != nil {
		t.Fatal(err1)
	}
	for i, _ := range got1.Ratings {
		if !proto.Equal(got1.Ratings[i], want1[i]) {
			t.Errorf("\n%v\n%v\n", got1.Ratings[i], want1[i])
		}
	}

	got2, err2 := client.GetRatings(ctx, &pb.GetRatingsRequest{ProductId: "N/A"})
	parsedRatings2, _ := parseRatings()
	want2 := parsedRatings2["N/A"]
	if err2 != nil {
		t.Fatal(err2)
	}
	for i, _ := range got2.Ratings {
		if !proto.Equal(got2.Ratings[i], want2[i]) {
			t.Errorf("\n%v\n%v\n", got2.Ratings[i], want2[i])
		}
	}
}
