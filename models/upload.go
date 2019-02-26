package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/uploadexpress/app/constants"
)

type Upload struct {
	Id    string  `json:"id" bson:"_id,omitempty"`
	Name  string  `json:"name" bson:"name"`
	Files []*File `json:"files" bson:"files"`
}

func (upload *Upload) BeforeCreate() {
	// Assigns an ID to each file
	for _, file := range upload.Files {
		if file.Id == "" {
			file.Id = bson.NewObjectId().Hex()
		}
	}

	upload.Id = bson.NewObjectId().Hex()
}

func (upload *Upload) Size() constants.ByteSize {
	var size int64 = 0
	for _, file := range upload.Files {
		size += file.Size
	}
	return constants.ByteSize(size)
}

const UploadsCollection = "uploads"