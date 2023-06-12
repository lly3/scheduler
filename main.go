package main

import (
	"context"
	"scheduler/interface/fiber"
	"scheduler/repository"
	usecase "scheduler/use_case"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := client.Database("scheduler")

	recordRepository := &repository.RecordRepo{
		Col: db.Collection("records"),
	}
	scheduleRepository := &repository.ScheduleRepo{
		Col: db.Collection("schedule"),
	}

	uc := &usecase.UseCase{
		RecordRepo: recordRepository,
		ScheduleRepo: scheduleRepository,
	}

	cfg := &fiber.ServerConfig{
		Port: "3000",
	}

	f := fiber.New(cfg, uc)

	f.Start(cfg)
}
