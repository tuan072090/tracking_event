package service

import (
	"context"
	"fmt"
	"time"
	"tracking_event/model"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func getClient() (*bigquery.Client, context.Context) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, GetProjectID())

	if err != nil {
		return nil, ctx
	}

	return client, ctx
}

// InsertEvent adds new event
func InsertEvent(event model.Event) error {

	client, ctx := getClient()

	u := client.Dataset(GetDatasetID()).Table(GetTableEvent()).Inserter()

	event.CreatedAt = int(time.Now().Unix())
	if err := u.Put(ctx, event); err != nil {
		return err
	}
	// [END bigquery_table_insert_rows]
	return nil
}

// QueryData runs raw query
func QueryData() ([]model.Event, error) {
	client, ctx := getClient()
	var arrResult []model.Event

	q := client.Query(
		"SELECT * FROM `meete-20160101-notification.meete_tracking.events` LIMIT 1000")
	// Location must match that of the dataset(s) referenced in the query.
	q.Location = "asia-southeast1"

	job, err := q.Run(ctx)
	if err != nil {
		return nil, err
	}

	status, err := job.Wait(ctx)
	if err != nil {
		return nil, err
	}
	if err := status.Err(); err != nil {
		return nil, err
	}
	it, err := job.Read(ctx)

	for {
		var row model.Event
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		fmt.Println("data row......")
		fmt.Println("----", row)
		fmt.Println("data row......")
		arrResult = append(arrResult, row)
	}

	countData("")
	return arrResult, err
}

func countData(query string) error {
	client, ctx := getClient()
	query = "SELECT count(Name) FROM `meete-20160101-notification.meete_tracking.events`"

	q := client.Query(query)
	// Location must match that of the dataset(s) referenced in the query.
	q.Location = "asia-southeast1"
	//q.DisableQueryCache = true

	job, err := q.Run(ctx)
	if err != nil {
		return err
	}

	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("data count is.......", row)
	}

	return nil
}

// CreateNewTable creates a new event table
func CreateNewTable() error {
	client, ctx := getClient()

	// bigquery.InferSchema infers BQ schema from native Go types.
	schema, err := bigquery.InferSchema(model.Event{})
	if err != nil {
		fmt.Println("err schema", err.Error())
		return err
	}

	table := client.Dataset(GetDatasetID()).Table(GetTableEvent())
	if err := table.Create(ctx, &bigquery.TableMetadata{Schema: schema}); err != nil {
		fmt.Println("err Create", err.Error())
		return err
	}

	return nil
}
