package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hcabnettek/filmapi/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	users  []*models.User
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("MDB_USER")
	pass := os.Getenv("MDB_PWORD")

	fmt.Println("Database connecting ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@bigid001-gyu7z.azure.mongodb.net/test?retryWrites=true&w=majority", user, pass)))
}
