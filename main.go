package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hcabnettek/filmapi/controllers"
	"github.com/hcabnettek/filmapi/repositories"
	"github.com/hcabnettek/filmapi/routers"
	"github.com/hcabnettek/filmapi/services"

	"github.com/joho/godotenv"
)

var (
	userRepository repositories.UserRepository = repositories.NewMongoRepository()
	userService    services.UserService        = services.NewUserService(userRepository)
	userController controllers.UserController  = controllers.NewUserController(userService)

	httpRouter routers.Router = routers.NewMuxRouter()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/users", userController.GetUsers)
	httpRouter.POST("/users", userController.AddUser)

	httpRouter.SERVE(os.Getenv("API_PORT"))
}
