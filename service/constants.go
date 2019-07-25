package service

import "os"

var projectID = "meete-20160101-notification"
var datasetID = "meete_tracking"
var tableEvent = "events"
var port = "8080"

// GetProjectID returns project id
func GetProjectID() string {
	return projectID
}

// GetDatasetID returns dataset id
func GetDatasetID() string {
	return datasetID
}

// GetTableEvent returns dataset id
func GetTableEvent() string {
	return tableEvent
}

// GetPort returns port number
func GetPort() string {
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	return port
}
