package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
/* MongoCN conexion a base de datos */
var MongoCN =  ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://ramiro:zi8hWGGZmCwiunVR@cluster0.bknkn.mongodb.net/twitter?authSource=admin&replicaSet=atlas-5dii1z-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true")

/* ConectarBD conexion a base de datos */
func ConectarBD() *mongo.Client  {
	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(),nil)
	if err !=nil{
		log.Fatal(err)
		return client
	}

	log.Println("Conexión Exitosa con la BD");
	return client 
}

/* "ChequeoConnection chequeo a base mediante ping" */
func ChequeoConnection() int {
	 err := MongoCN.Ping(context.TODO(),nil)
	 if err != nil {
		 return 0
	 }
	 return 1
}