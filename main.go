package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"net"
)

func main() {
	var ip string
	flag.StringVar(&ip, "ip", "127.0.0.1", "get ip of mongo db")
	flag.Parse()

	if net.ParseIP(ip) == nil {
		panic(fmt.Errorf("invalid ip address %s", ip))
	}

	var uri = fmt.Sprintf("mongodb://%s:27017", ip)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println("Can not Connect to mongodb")
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		fmt.Printf("Unable to Ping MongoDB at %s\n", uri)
		panic(err)
	}
	fmt.Printf("Pinged. Successfully connected to MongoDB! at %s\n", uri)

}
