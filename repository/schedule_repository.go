package repository

import (
	"context"
	"scheduler/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScheduleRepo struct{
	Col *mongo.Collection
}

type ScheduleDoc struct{
	Id string `bson:"id"`
	Todos []ScheduleItemDoc `bson:"todos"`
}

type ScheduleItemDoc struct{
	Title string `bson:"title"`
	Duration time.Duration `bson:"duration"`
}

func (s ScheduleRepo) GetScheduleById(scheduleId string) (entities.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var scheduleDoc ScheduleDoc
	err := s.Col.FindOne(ctx, bson.D{{"id",scheduleId}}).Decode(&scheduleDoc)
	if err != nil {
		return entities.Schedule{}, err
	}

	return scheduleDocToSchedule(scheduleDoc), nil
}

func (s ScheduleRepo) GetAllSchedule() ([]entities.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := s.Col.Find(ctx, bson.D{{}})
	if err != nil {
		return []entities.Schedule{}, err
	}
	var scheduleDoc []ScheduleDoc
	
	if err := res.All(context.Background(), &scheduleDoc); err != nil {
		return []entities.Schedule{}, err
	}

	var schedule []entities.Schedule
	for _, v := range scheduleDoc {
		schedule = append(schedule, scheduleDocToSchedule(v))
	}

	return schedule, nil
}

func scheduleDocToSchedule(scheduleDoc ScheduleDoc) entities.Schedule {
	var scheduleItems []entities.ScheduleItem
	for _, item := range scheduleDoc.Todos {
		scheduleItems = append(scheduleItems, entities.ScheduleItem{
			Title: item.Title,
			Duration: item.Duration,
		})
	}

	return entities.Schedule{
		Id: scheduleDoc.Id,
		Todos: scheduleItems,
	}
}
