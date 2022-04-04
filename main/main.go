package main
import (
    "net/http"
	"time"
	"context"
	"fmt"
    "github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/Alejandraarrieta/atenthication/database"
	"github.com/Alejandraarrieta/atenthication/models"
	"github.com/Alejandraarrieta/atenthication/jwt"
)
var client *mongo.Client

func main(){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongoConnect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/login", CreateUser).Methods("POST")
	http.ListenAndServe(":8080",router)
	
}