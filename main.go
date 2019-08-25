package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/nayanmakasare/transavro/handler"
	"github.com/nayanmakasare/transavro/subscriber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"

	transavro "github.com/nayanmakasare/transavro/proto/transavro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("cloudwalker.srv.transavro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	uri := os.Getenv("DB_HOST")
	log.Log(uri)
	if uri == "" {
		log.Log("empty mila")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"));
	if err != nil {
		log.Log(err)
	}

	collection := client.Database("test").Collection("movies")

	// Register Handler
	transavro.RegisterTransavroHandler(service.Server(), &handler.Transavro{collection})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("cloudwalker.srv.transavro", service.Server(), new(subscriber.Transavro))

	// Register Function as Subscriber
	micro.RegisterSubscriber("cloudwalker.srv.transavro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
