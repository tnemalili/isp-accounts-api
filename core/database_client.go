package core

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

// CONNECTION TO DATABASE

var ctx = context.Background()

func DatabaseConnection(config *DatabaseConfig) ( *mongo.Database, error) {

	// SETTING LOGGER
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceColors: true})

	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	dbClient, err := mongo.NewClient(clientOptions)

	if err != nil { log.Info("Mongodb Error on Creating New to mongodb client") }

	err = dbClient.Connect(ctx)

	if err != nil { log.Info("Mongodb Error on Connecting to mongodb")}

	// Ping Database Client Here

	ping := dbClient.Ping(ctx, readpref.Primary())

	if ping != nil { log.Info("Pong: couldn't  connect to mongodb") }

	log.Info("&&&&&&&&&&&&&&&&&&&&&&&& Successfully connected to Database &&&&&&&&&&&&&&&&&&&&&&&&")

	return dbClient.Database(config.DatabaseName), nil
}

var (
	hostname = os.Getenv("DB_HOST")
	name     = os.Getenv("DB_NAME")
	port     = os.Getenv("DB_PORT")
	URI = fmt.Sprintf(`mongodb://%s`, hostname)
)

// DATABASE CONNECTION CONFIG STRUCT

type DatabaseConfig struct {
	ConnectionString	string	`json:"connection_string"`
	Enabled				bool	`json:"enabled"`
	Port				string	`json:"port"`
	DatabaseName		string	`json:"database_name"`
}

// DATABASE CONNECTION CONFIG

var DBConnConfig = DatabaseConfig{
	ConnectionString: URI,
	Enabled:          true,
	Port:             port,
	DatabaseName:     name,
}

// DATABASE  CONNECTION
var Database, _ = DatabaseConnection(&DBConnConfig)

// REPOSITORY  STRUCT TO BE ACCESSIBLE CODE WIDE

type Repository struct { Database *mongo.Database }


// REPOSITORY INSTANCE

var DBClient = &Repository{ Database: Database}