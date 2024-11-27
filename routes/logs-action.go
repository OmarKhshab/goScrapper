package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"

	"example.com/rest-api/mongoose"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/api/option"
)

func SaveLogs() {
	var db = mongoose.MongoDB
	id := uuid.New()
	var ctx = context.TODO()
	var logsCollection = db.Collection("logs")
	query, err := logsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("bad query")
	}
	type logInterface struct {
		URL    string `bson:"url" json:"url"`
		Method string `bson:"method" json:"method"`
		Status int32  `bson:"status" json:"status"`
	}
	var logs = make([]logInterface, 0)
	err = query.All(ctx, &logs)
	if err != nil {
		log.Fatal("bad server ", err)
	}
	defer query.Close(ctx)
	opt := option.WithCredentialsFile("omarGolang.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Storage(ctx)
	if err != nil {
		fmt.Errorf("error in app: %v", err)
	}
	handler, err := client.Bucket("erp-system-1a749.appspot.com")
	if err != nil {
		fmt.Errorf("error in bucket: %v", err)
	}
	handlerone := handler.Object("logss.json")
	Writer := handlerone.NewWriter(context.Background())
	Writer.ContentType = "application/json"
	Writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	logsBytes, err := json.Marshal(logs)
	if err != nil {
		fmt.Errorf("error in parsing: %v", err)
	}
	_, errs := Writer.Write(logsBytes)
	if errs != nil {
		fmt.Errorf("error in writing: %v", errs)
	}
	err = Writer.Close()
	if errs != nil {
		fmt.Errorf("error in closing: %v", errs)
	}
}
