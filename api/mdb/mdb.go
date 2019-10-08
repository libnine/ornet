package mdb

import (
	"context"
	"log"
	"os"
	"time"

	models "../models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("ORNET")))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func FindAllCases(client *mongo.Client) []models.Case {
	var cases []models.Case

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := client.Database("ornet").Collection("cases").Find(ctx, bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "case_date", Value: 1}}))
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var d models.Case
		err = cur.Decode(&d)
		cases = append(cases, d)
	}

	return cases
}

func FindCasesByUser(client *mongo.Client, user string) []models.Case {
	var cases []models.Case

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := client.Database("ornet").Collection("cases").Find(ctx, bson.D{primitive.E{Key: "rep", Value: user}}, options.Find().SetSort(bson.D{primitive.E{Key: "case_date", Value: 1}}))
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var d models.Case
		err = cur.Decode(&d)
		cases = append(cases, d)
	}

	return cases
}

func FindAllCasesHistorical(client *mongo.Client) []models.Case {
	var cases []models.Case

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := client.Database("ornet").Collection("cases_historical").Find(ctx, bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "case_date", Value: -1}}))
	if err != nil {
		log.Fatal(err)
	}

	return cases
}

func FindAllEq(client *mongo.Client) []models.Eq {
	var eq []models.Eq

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := client.Database("ornet").Collection("equipment").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var d models.Eq
		err = cur.Decode(&d)
		eq = append(eq, d)
	}

	return eq
}

func FindAllUsers(client *mongo.Client) []models.User {
	var users []models.User

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := client.Database("ornet").Collection("users").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var d models.User
		err = cur.Decode(&d)
		users = append(users, d)
	}

	return users
}

func InsertCase(client *mongo.Client, _case *models.Case) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ins, err := client.Database("ornet").Collection("cases").InsertOne(ctx, _case)
	if err != nil {
		log.Fatal(err)
	}

	return ins.InsertedID
}

func InsertUser(client *mongo.Client, _user *models.User) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ins, err := client.Database("ornet").Collection("users").InsertOne(ctx, _user)
	if err != nil {
		log.Fatal(err)
	}

	return ins.InsertedID
}
