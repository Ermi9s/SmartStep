package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Location    string             `json:"location,omitempty" bson:"location,omitempty"`
	Status string             `json:"status,omitempty" bson:"status,omitempty"`
}

type LocationData struct {
	Name string
	NumberOfDevices int32
}