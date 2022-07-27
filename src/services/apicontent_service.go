package services

import (
	apicontent "github.com/henrikroschmann/F-api-ker/domain/api-content"
	"github.com/henrikroschmann/F-api-ker/utils/errors"
)

func CreateApiContent(jsons apicontent.Jsons) (*apicontent.Jsons, *errors.RestErr) {

	if err := jsons.Save(); err != nil {
		return nil, err
	}

	return &jsons, nil
}

func GetApiContent() (*apicontent.JsonsSlice, *errors.RestErr) {
	result := &apicontent.JsonsSlice{}

	data, err := result.Get()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetApiContentById(id string) (*apicontent.Jsons, *errors.RestErr) {
	result := &apicontent.Jsons{ContentId: id}

	if err := result.GetById(id); err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteApiContent(contentID string) *errors.RestErr {
	contRemove := &apicontent.Jsons{}
	if err := contRemove.DeleteById(contentID); err != nil {
		return err
	}

	return nil
}
