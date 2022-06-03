package service

import (
	"context"
	"fmt"
	"github.com/jszwec/csvutil"
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	consumer "github.com/martinkaburu/warehouse-manager/warehouseproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"strconv"
	"strings"
)

func Stream(orders *csvutil.Decoder, port *int) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), opts...)
	if err != nil {
		log.Fatalf("Failed to connect to remote server: %v", err)
	}

	defer conn.Close()

	client := consumer.NewOrderConsumerClient(conn)
	var entry models.OrderEntry

	log.Printf("Streaming Orders to Server")
	ctx := context.Background()

	stream, err := client.ReceiveOrders(ctx)
	for {
		if err := orders.Decode(&entry); err == io.EOF {
			stream.CloseSend()
			break
		}
		weight, _ := strconv.ParseFloat(strings.TrimSpace(entry.Weight), 64)

		order := consumer.Order{
			Weight: weight,
			Email:  entry.Email,
			Id:     entry.Id,
			Phone:  entry.Phone,
		}
		if err := stream.Send(&consumer.OrderRequest{Order: &order}); err != nil {
			log.Fatalf("%v.Send(%v) = %v: ", stream, entry, err)
		}
	}

	log.Printf("Done Streaming Orders to Server")
}
