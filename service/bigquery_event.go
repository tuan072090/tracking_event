package service

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"time"
	"tracking_event/model"
)

func GetClient() (*bigquery.Client, context.Context) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, ProjectId)

	if err != nil {
		return nil, ctx
	}

	return client, ctx
}

func InsertEvent(event model.Event) error {

	client, ctx := GetClient()

	u := client.Dataset(DatasetId).Table(TableEvent).Inserter()

	event.CreatedAt = int(time.Now().Unix())
	if err := u.Put(ctx, event); err != nil {
		return err
	}
	// [END bigquery_table_insert_rows]
	return nil
}

// CreateTable creates a new event table
func CreateNewTable() error {
	client, ctx := GetClient()

	// bigquery.InferSchema infers BQ schema from native Go types.
	schema, err := bigquery.InferSchema(model.Event{})
	if err != nil {
		fmt.Println("err schema", err.Error())
		return err
	}

	table := client.Dataset(DatasetId).Table(TableEvent)
	if err := table.Create(ctx, &bigquery.TableMetadata{Schema: schema}); err != nil {
		fmt.Println("err Create", err.Error())
		return err
	}

	return nil
}
