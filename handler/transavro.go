package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/micro/go-micro/util/log"

	transavro "github.com/nayanmakasare/transavro/proto/transavro"
)

type Transavro struct{
	MyMongoCollection *mongo.Collection
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Transavro) Call(ctx context.Context, req *transavro.Request, rsp *transavro.Response) error {
	log.Log("Received Transavro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Transavro) Stream(ctx context.Context, req *transavro.StreamingRequest, stream transavro.Transavro_StreamStream) error {
	log.Logf("Received Transavro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&transavro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}
	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Transavro) PingPong(ctx context.Context, stream transavro.Transavro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&transavro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (e *Transavro) GetTiles(ctx context.Context, req *transavro.RequestSpecification, stream transavro.Transavro_GetTilesStream) error {
	cur, err := e.MyMongoCollection.Find(ctx, bson.D{{}}, nil)
	if err != nil {
		return err
	}
	for cur.Next(ctx){
		var result *transavro.MovieTiles
		err = cur.Decode(result)
		if err != nil {
			return err
		}
		err = stream.Send(result)
		if err != nil {
			return err
		}
	}

	err = cur.Close(ctx)
	if err != nil {
		return err
	}
	return nil
}








































