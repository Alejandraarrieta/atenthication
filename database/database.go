package database
import(
	"context"
	"encoding/json"
	"net/http"
	"time"
	//"go.mongodb.org/mongo-driver/mongo" //imported and not used:
	//"go.mongodb.org/mongo-driver/bson/primitive" //imported and not used:
	"github.com/Alejandraarrieta/atenthication/models"
	//"github.com/mongodb/mongo-go-driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//funcion que crea e inserta user database
func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	var user UserData
	_ = json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("UserDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _:= collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
}

//func que busca un usuario
func SearchUser(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	var users [] UserData
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
