package main
import (
    "net/http"
	"encoding/json"
	"time"
	"context"
	"fmt"
    "github.com/gorilla/mux"
	//"github.com/mongodb/mongo-go-driver/mongo"
	//"github.com/Alejandraarrieta/atenthication/database"
	//"github.com/Alejandraarrieta/atenthication/models"
	//"github.com/Alejandraarrieta/atenthication/jwt"
	// "github.com/mongodb/mongo-go-driver"
	//"go.mongodb.org/mongo-driver"
	"github.com/Alejandraarrieta/atenthication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	 
)
var client *mongo.Client

func main(){
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/login", CreateUser).Methods("POST")
	http.ListenAndServe(":8080",router)
	
}

//funcion que crea e inserta user database
func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("UserDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _:= collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
}

//func que busca un usuario
func SearchUser(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	var users [] User
	collection := client.Database("UserDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error()+ `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
	var user User
	cursor.Decode(&user)
	users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error()+ `"}`))
		return
	}
	json.NewEncoder(w).Encode(users)

}