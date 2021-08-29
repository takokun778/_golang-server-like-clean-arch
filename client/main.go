package main

import (
	"context"
	"log"
	"template/proto/template"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.TODO()

	conn, err := grpc.Dial("localhost:19003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	templateClient := template.NewTemplateServiceClient(conn)

	templateDeleteRequest := new(template.DeleteRequest)
	id, _ := uuid.NewRandom()
	templateDeleteRequest.Id = id.String()

	_, err = templateClient.Delete(ctx, templateDeleteRequest)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println(st.Code())
		log.Println(st.Err())
		log.Println(st.Message())
		log.Println(st.Proto())
		for _, detail := range st.Details() {
			log.Println(detail)
		}
	}
}
