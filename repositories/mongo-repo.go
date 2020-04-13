package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hcabnettek/filmapi/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct{}

// NewMongoRepository constructor function
func NewMongoRepository() UserRepository {
	return &repo{}
}

func (*repo) FindAll() ([]models.User, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@bigid001-gyu7z.azure.mongodb.net/test?retryWrites=true&w=majority", os.Getenv("MDB_USER"), os.Getenv("MDB_PWORD"))))
	if err != nil {
		log.Fatal("Connection issues")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	defer cancel()

	var users []models.User
	movieDatabase := client.Database("sample_mflix")
	usersCollection := movieDatabase.Collection("users")

	cursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return users, err
	}

	return users, nil
}

func (*repo) Save(user *models.User) (*models.User, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@bigid001-gyu7z.azure.mongodb.net/test?retryWrites=true&w=majority", os.Getenv("MDB_USER"), os.Getenv("MDB_PWORD"))))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	defer cancel()

	return nil, nil
}
