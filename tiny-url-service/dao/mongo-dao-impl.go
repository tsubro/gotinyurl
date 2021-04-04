package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/bson"
	"tiny-url/models"
	"tiny-url/utils"
)

func Insert(r *models.UrlModel) {

	collection := utils.GetMongoConnection().Database("tinydb").Collection("url-test")

	insertResult, err := collection.InsertOne(context.TODO(), r)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Inserted a single document: ", insertResult.InsertedID)
}

// func Get(jobId string)  {
// 	collection := utils.GetMongoConnection().Database("document-renderer").Collection("render-job")

// 	var job models.RenderJob

// 	err := collection.FindOne(context.TODO(), bson.D{{"jobid", jobId}}).Decode(&job)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	log.Info("Fetched a single document succefully")
// 	return &job, nil
// }

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