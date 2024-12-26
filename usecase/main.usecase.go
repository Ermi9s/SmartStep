package usecase

import (
	"context"
	"step-project/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Main_Usecase struct {
	Collection *mongo.Collection
}

func New_Usecase(cl *mongo.Collection) *Main_Usecase {
	return &Main_Usecase{
		Collection: cl,
	}
}

func (mu *Main_Usecase)GetAllInfo()([]domain.LocationData , error) {
	var data []domain.LocationData
	
	cursor,err :=  mu.Collection.Find(context.TODO() , bson.D{})
	if err != nil {
		return nil , err
	}


	mapping := make(map[string]int32)
	for cursor.Next(context.TODO()) {
		var device domain.Device

		if err := cursor.Decode(&device); err != nil {
			return nil, err
		}

		if device.Status == "active" {
			mapping[device.Location]++
		}
	}

	for key,val := range mapping {
		var single domain.LocationData
		single.Name = key
		single.NumberOfDevices = val

		data = append(data, single)
	}

	return data,nil

}


func (mu *Main_Usecase) RegisterDevice(device domain.Device) error {
    _, err := mu.Collection.InsertOne(context.TODO(), device)
    if err != nil {
        return err
    }
    return nil
}

func (mu *Main_Usecase) UpdateStatus(device domain.Device, status string) error {
	filter := bson.D{{Key : "_id", Value:  device.ID}}
	update := bson.D{{Key : "$set", Value: bson.D{{Key : "status", Value: status},}},}

	_, err := mu.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
