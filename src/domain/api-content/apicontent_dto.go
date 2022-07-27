package apicontent

import "go.mongodb.org/mongo-driver/bson/primitive"

type Jsons struct {
	ID        primitive.ObjectID     `bson:"_id" json:"-"`
	Content   map[string]interface{} `json:"content"`
	ContentId string                 `json:"content_id"`
}

type JsonsSlice []Jsons
