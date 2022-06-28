package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/milinches/edsu-token-backend/api/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("error disconnecting: %v", err)
		}
	}()
	
	log.Println("Connected to mongo successfully...")

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("ping error: %v", err)
	}

	r := router.NEW()

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}