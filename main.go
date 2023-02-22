package main


import (
    "github.com/url/shorter/handlers/router"
	"github.com/url/shorter/database/redis"
	"log"
	"os"
)


func main(){
	database.InitRedis()

	log.Fatal(router.APIRouter.Run(":"+os.Getenv("PORT")))
}