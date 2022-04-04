package models

import ( 
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/mongodb/mongo-go-driver"
)
//struct to create token with the user and claims standard
type Claim struct{
	User  `json:"user"`
	jwt.StandardClaims
}

//struct to return token 
type ResponseToken struct{
	Token string `json:"token"`
}

type User struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Lastname string `json:"lastname" bson:"lastname,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Photo    string `json:"photo,omitempty" bson:"photo,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Workunit string `json:"worjunit,omitempty" bson:"lastname,omitempty"`
	Area     string `json:"area,omitempty" bson:"area,omitempty"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
	//CreatedAt time.Time `json:"createdat,omitempty""` //fecha creacion
	//UpdatedAt time.Time `json:"updatedat,omitempty""` //fecha actualizacion
}