package apicontent

import (
	"context"
	"log"
	"time"

	apicontentdb "github.com/henrikroschmann/F-api-ker/datasource/mongo/apiContent_db"
	"github.com/henrikroschmann/F-api-ker/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (jsons *Jsons) Save() *errors.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	jsons.ID = primitive.NewObjectID()
	jsons.ContentId = jsons.ID.Hex()

	if _, err := apicontentdb.Collection.InsertOne(ctx, jsons); err != nil {
		defer cancel()
		return errors.NewInternalServerError("content could not be saved")
	}
	defer cancel()

	return nil
}

func (jsons *JsonsSlice) Get() (*JsonsSlice, *errors.RestErr) {
	var jSlice JsonsSlice
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := apicontentdb.Collection.Find(ctx, bson.M{})
	if err != nil {
		defer cancel()
		return nil, errors.NewInternalServerError("No content available")
	}

	for result.Next(ctx) {

		var elem Jsons
		err := result.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		jSlice = append(jSlice, elem)
	}

	defer cancel()
	return &jSlice, nil
}

func (jsons *Jsons) GetById(id string) *errors.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := apicontentdb.Collection.FindOne(ctx, bson.M{"contentid": id}).Decode(&jsons)
	defer cancel()
	if err != nil {
		return errors.NewInternalServerError("invalid contentid")
	}

	return nil
}

func (jsons *Jsons) DeleteById(contentID string) *errors.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	_, err := apicontentdb.Collection.DeleteOne(ctx, bson.M{"contentid": contentID})
	defer cancel()
	if err != nil {
		return errors.NewInternalServerError("invalid contentid")
	}
	return nil
}
