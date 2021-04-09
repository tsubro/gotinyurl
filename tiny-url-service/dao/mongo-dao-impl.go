package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"tiny-url/models"
	"tiny-url/utils"
)

func Insert(r *models.UrlModel) error  {

	collection := utils.GetMongoConnection().Database("tinydb").Collection("url")

	insertResult, err := collection.InsertOne(context.TODO(), r)
	if err != nil {
		return err
	}

	log.Info("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func Get(hash string) (*models.UrlModel, error) {
	collection := utils.GetMongoConnection().Database("tinydb").Collection("url")

	var r models.UrlModel

	err := collection.FindOne(context.TODO(), bson.D{{"hash", hash}}).Decode(&r)
	if err != nil {
		return nil, err
	}
	log.Info("Fetched a single document succefully")
	return &r, nil
}

// func Update() {
// 	collection := utils.GetMongoConnection().Database("document-renderer").Collection("render-job")

// 	update := bson.M{
// 		"$set": bson.M{
// 		  "state": job.State,
// 		  "inputs": job.Inputs,
// 		},
// 	  }
// 	_, err := collection.UpdateOne(context.TODO(), bson.D{{"jobid", job.JobId}}, update)
// 	if err != nil {
// 		log.Error(err)
// 	}
// }