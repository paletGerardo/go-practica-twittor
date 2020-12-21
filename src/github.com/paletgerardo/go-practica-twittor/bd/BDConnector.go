package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCn = connectToMongoDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://gerrDev:120383..@cluster0.juukd.mongodb.net/<dbname>?retryWrites=true&w=majority")

/* Conectar a BD */
func connectToMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions )

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conectado...")
	return client
}

/* Ping a la BD */
func CheckConection() int {
	err := MongoCn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
