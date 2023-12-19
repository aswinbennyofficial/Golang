package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key
var JWT_SECRET_KEY = []byte("very-secret-key")

// map of username and password
var credMap = map[string]string{
	"user":"pass",
	"aswinbenny": "password123",
	"willywonka": "chocolatefactory",
}

// Struct for unmarshalling from form and other uses
type Credentials struct{
	Username string `json:"username"`
	Password string `json:"password"`
}


// Claims for jwt
type Claims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims

}


func main(){
	http.HandleFunc("/login",loginHandler)
	http.HandleFunc("/home",homeHandler)
	http.HandleFunc("/refresh",refreshTokenHandler)


	fmt.Println("Starting server in port 8080")
	err:=http.ListenAndServe(":8080",nil)

	if err!=nil{
		log.Printf("Error starting the server at port 8080..")
	}
}


// handlers
func loginHandler(w http.ResponseWriter, r *http.Request) {
	
	var creds Credentials
	err:=json.NewDecoder(r.Body).Decode(&creds)

	if err!=nil{
		log.Println("ERROR DECODING JSON ",err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("USERNAME:%s  PASSWORD;%s \n",creds.Username,creds.Password)

	// Check if password is correct
	if creds.Password==credMap[creds.Username] {
		log.Println("PASSWORD IS CORRECT")
	} else{
		log.Println("INCORRECT PASSWORD")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// setting JWT claims
	// Setting experiation time to be 15 minutes from now
	expirationTime :=time.Now().Add(5 * time.Minute)

	// Creating payload
	claims:= &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declaring token with header and payload
	noSignedToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	// Create a complete signed JWT
	signedToken,err:=noSignedToken.SignedString(JWT_SECRET_KEY)

	if err!=nil{
		log.Println("ERROR OCCURED WHILE CREATING JWT TOKEN: ",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,&http.Cookie{
		Name: "JWtoken",
		Value: signedToken,
		Expires: expirationTime,
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	// Takes the cookie from the request named JWtoken
	cookieVar, err:=r.Cookie("JWtoken")
	if err!=nil{
		log.Println("ERROR WHILE OBTAINING COOKIE: ",err)
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Get JWT string from cookie
	JWTstring:=cookieVar.Value

	// Initialising a new instance of claims
	claims:=&Claims{}

	// Parse the JWT string and store in claims
	tokenVar,err:= jwt.ParseWithClaims(JWTstring,claims,func(token *jwt.Token) (any, error){
		return JWT_SECRET_KEY,nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("INVALID SIGNATURE")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("ERROR WHILE PARSING JWT WITH CLAIMS: ",err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tokenVar.Valid {
		log.Println("INVALID TOKEN ")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request){
	
}