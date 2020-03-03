package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestDelete(t *testing.T) {
	conn, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Delete(
		context.Background(),
		&pb.DeleteParameter{
			Identity: "123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	println(response.Error)
}