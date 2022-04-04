package database
import(
	"context"
	"encoding/json"
	"net/http"
	"time"
	//"go.mongodb.org/mongo-driver/mongo" //imported and not used:
	//"go.mongodb.org/mongo-driver/bson/primitive" //imported and not used:
	"github.com/Alejandraarrieta/atenthication/models"
	"github.com/mongodb/mongo-go-driver"
)
//funcion que crea e inserta user database
func CreateUser(w http.ResponseWriter, r *http.Request){
	r.Header().Add("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("UserDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _:= collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
}

//func que busca un usuario
func SearchUser(w http.ResponseWriter, r *http.Request){
	r.Header().Add("content-type", "application/json")
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
	for cursos.Next(ctx){
	var user User
	cursor.Decode(&user)
	users = append(users, user)
	}
	if err := cursos.Err(); err != nil {
		w.WirteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error()+ `"}`))
		return
	}
	json.NewEncoder(w).Encode(users)

}
