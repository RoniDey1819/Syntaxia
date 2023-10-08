package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	model "github.com/PratikforCoding/Syntaxia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIConfig struct {
	AttendeesCollection *mongo.Collection
}

func NewAPIConfig(attendeesCol *mongo.Collection) *APIConfig {
	return &APIConfig{AttendeesCollection: attendeesCol}
}

func (apiCfg *APIConfig) register(attendee model.Attendee) (model.Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Generate a user ID (e.g., a unique serial number)
	count, err := apiCfg.AttendeesCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	countStr := strconv.Itoa(int(count) + 1)
	attendee.SerialNo = countStr

	// Set other attendee properties
	attendee.Taken = "no"

	// Insert the attendee into the database
	inserted, err := apiCfg.AttendeesCollection.InsertOne(context.Background(), attendee)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted user id:", inserted.InsertedID)

	// Return the created attendee
	return attendee, nil
}

func (apiCfg *APIConfig) getAttendee(serial string) (model.Attendee, error) {
	filter := bson.M{"serialno": serial}
	var user model.Attendee
	err := apiCfg.AttendeesCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("user not found")
			return model.Attendee{}, errors.New("user not found")
		} else {
			log.Fatal(err)
		}
	}

	return user, nil
}

func (apiCfg *APIConfig) claimefood(serial string) (model.Attendee, error) {
	filter := bson.M{"serialno": serial}
	update := bson.M{"$set": bson.M{"taken": "yes"}}

	updateResult, err := apiCfg.AttendeesCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return model.Attendee{}, err
	}

	if updateResult.ModifiedCount == 0 {
		return model.Attendee{}, errors.New("no document was updated")
	}

	var updatedAttendee model.Attendee
	err = apiCfg.AttendeesCollection.FindOne(context.TODO(), filter).Decode(&updatedAttendee)
	if err != nil {
		return model.Attendee{}, err
	}

	return updatedAttendee, nil
}

func (apiCfg *APIConfig) getAllAttendees() ([]model.Attendee, error) {
	var attendees []model.Attendee
	filter := bson.M{}

	cursor, err := apiCfg.AttendeesCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var attendee model.Attendee
		if err := cursor.Decode(&attendee); err != nil {
			return nil, err
		}
		attendees = append(attendees, attendee)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return attendees, nil
}

func (apiCfg *APIConfig) getAttendeesByYear(year string) ([]model.Attendee, error) {
	var attendees []model.Attendee

	filter := bson.M{"year": year}

	cursor, err := apiCfg.AttendeesCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var attendee model.Attendee
		if err := cursor.Decode(&attendee); err != nil {
			return nil, err
		}
		attendees = append(attendees, attendee)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return attendees, nil
}
