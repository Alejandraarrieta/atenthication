package jwt

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"time"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/Alejandraarrieta/atenthication/database"
	"github.com/Alejandraarrieta/atenthication/models"
)

var (
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
)

func init(){
	privateBytes, err := ioutil.ReadFile("./private.rsa") //crear key y pasar ruta del archivo
	if err != nil{
		log.Fatal("could not read the file")
	}
	publicBytes, err := ioutil.ReadFile("./public.rsa.pub") //crear key y pasar ruta del archivo
	if err != nil{
		log.Fatal("could not read the file")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil{
		log.Fatal("Could not parse privatekey")
	}
	publicKey, err = jwt.ParseRSAPublickeyFromPEM(publicBytes)
	if err != nil{
		log.Fatal("Could not parse publickey")
	}

}

func GenerateJWT(user User) (string, err){
	claims := models.Claim{
		User : user,
		StandardClaims:jwt.StandardClaims{
			ExpireAt: time.Now().Add(time.Hour * 1).Unix(),
		}
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil{
		log.Fatal("No se pudo firmar")
	}
	return result, nil
}


func Login(w http.ResponseWriter, r *http.Request){	
	user := SearchUser(w, r)
	token := GenerateJWT(user)
	result := models.ResponseToken{token}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Fprinln(w, "Error al generar json")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}




func ValidateToken(w http.ResponseWriter, r *http.Request){
	token, err:= request.ParseFromRequestWhitClaims(r, request.OAuth2Extractor, &models.claims{},func(token *jwt.Token)(interface{}, error){
		return publickey, nil
	})
	if err != nil{
		switch err.(type){
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Fprinln(w, "Expired token")
				return 
			case jwr.ValidationErrorSignatureInvallid:
				fmt.Fprintln(w, "The token signature is invalid")
				return
			default:
				fmt.Fprintln(w, "The token is not valid")
				return
			}
		//Si no es error de validacion:
		default:
			fmt.Fprintln(w, "The token is not valid")
				return
		}
	}
	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprinln(w,"Welcome to the system")
	}else{
		w.WirteHeader(http.StatusUnauthorized)
		fmt.Fprinln(w, "Your token is not valid")
	}	
}











/*func getUser(c *gin.Context) {
    email := c.Param("email")
    for _, a := range users {
        if a.Email == email {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
//Tengo que recibir la peticion y leer el cuerpo
func resposeJWT(w http.ResponseWriter, r *http.Request){
	var username string
	err :=jsonNewDecoder(r.Body).Decode(&username)
	token := generataJWT(getUser(username))
	result := models.ResponseToken{token}
	jsonResul, err := json.Marshal(result)
	if err != nil{
		fmt.Fprintln(w, "Error al generar el json")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
	
}*/