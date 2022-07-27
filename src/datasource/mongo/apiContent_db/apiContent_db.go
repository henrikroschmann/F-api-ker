package apicontentdb

import (
	database "github.com/henrikroschmann/F-api-ker/datasource/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection = database.OpenCollection(database.Client, "apicontent")
