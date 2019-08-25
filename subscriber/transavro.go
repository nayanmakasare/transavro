package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	transavro "github.com/nayanmakasare/transavro/proto/transavro"
)

type Transavro struct{}

func (e *Transavro) Handle(ctx context.Context, msg *transavro.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *transavro.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
