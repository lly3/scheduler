package repository

import (
	"context"
	"scheduler/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordDoc struct {
	Id string `bson:"id"`
	ScheduleId string `bson:"schedule_id"`
	Items []RecordItemDoc `bson:"items"`
}

type RecordItemDoc struct {
	Title string `bson:"title"`
	Start time.Time `bson:"start"`
	End time.Time `bson:"end"`
}

type RecordRepo struct{
	Col *mongo.Collection
}

func (r RecordRepo) GetRecordById(recordId string) (entities.Record, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var recordDoc RecordDoc
	err := r.Col.FindOne(ctx, bson.D{{"id", recordId}}).Decode(&recordDoc)
	if err != nil {
		return entities.Record{}, err
	}

	return recordDocToRecord(recordDoc), nil
}

func (r RecordRepo) Insert(record entities.Record) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var recordDoc RecordDoc
	recordDoc = recordToRecordDoc(record)
	_, err := r.Col.InsertOne(ctx, recordDoc)
	if err != nil {
		return err
	}

	return nil
}

func (r RecordRepo) Update(record entities.Record) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var recordDoc RecordDoc
	recordDoc = recordToRecordDoc(record)
	_, err := r.Col.UpdateOne(ctx, bson.D{{"id",record.Id}}, bson.D{{"$set", recordDoc}})
	if err != nil {
		return err
	}

	return nil
}

func recordDocToRecord(recordDoc RecordDoc) entities.Record {
	var recordItems []entities.RecordItem

	for _, item := range recordDoc.Items {
		recordItems = append(recordItems, entities.RecordItem{
			Title: item.Title,
			Start: item.Start,
			End: item.End,
		})
	}

	return entities.Record{
		Id: recordDoc.Id,
		ScheduleId: recordDoc.ScheduleId,
		Items: recordItems,
	} 
}

func recordToRecordDoc(record entities.Record) RecordDoc {
	var recordItemsDoc []RecordItemDoc

	for _, item := range record.Items {
		recordItemsDoc = append(recordItemsDoc, RecordItemDoc{
			Title: item.Title,
			Start: item.Start,
			End: item.End,
		})
	}

	return RecordDoc{
		Id: record.Id,
		ScheduleId: record.ScheduleId,
		Items: recordItemsDoc,
	}
}
