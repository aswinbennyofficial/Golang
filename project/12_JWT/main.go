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
	"user":        "pass",
	"aswinbenny":  "password123",
	"willywonka":  "chocolatefactory",
}

// Struct for unmarshalling from form and other uses
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims for jwt
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Generate a new JWT token
func generateToken(username string) (string, error) {
	// Setting expiration time to be 5 minutes from now
	expirationTime := time.Now().Add(5 * time.Minute)

	// Creating payload
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declaring token with header and payload
	noSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create a complete signed JWT
	signedToken, err := noSignedToken.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Parse and validate JWT from request
func parseAndValidateJWT(r *http.Request) (*Claims, error) {
	cookieVar, err := r.Cookie("JWtoken")
	if err != nil {
		return nil, err
	}

	// Get JWT string from cookie
	JWTstring := cookieVar.Value

	// Initialize a new instance of Claims
	claims := &Claims{}

	// Parse the JWT string and store in claims
	tokenVar, err := jwt.ParseWithClaims(JWTstring, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenVar.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/refresh", refreshTokenHandler)

	fmt.Println("Starting server in port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Printf("Error starting the server at port 8080..")
	}
}

// Handlers

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		log.Println("ERROR DECODING JSON ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("USERNAME:%s  PASSWORD:%s \n", creds.Username, creds.Password)

	// Check if password is correct
	if creds.Password == credMap[creds.Username] {
		log.Println("PASSWORD IS CORRECT")
	} else {
		log.Println("INCORRECT PASSWORD")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new JWT token
	signedToken, err := generateToken(creds.Username)
	if err != nil {
		log.Println("ERROR OCCURRED WHILE CREATING JWT TOKEN: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Setting JWT claims
	expirationTime := time.Now().Add(5 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "JWtoken",
		Value:   signedToken,
		Expires: expirationTime,
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and validate JWT from request
	claims, err := parseAndValidateJWT(r)
	if err != nil {
		log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("ACCESS APPROVED TO /home")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and validate JWT from request
	claims, err := parseAndValidateJWT(r)
	if err != nil {
		log.Println("ERROR WHILE PARSING/VALIDATING JWT: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// New token is not issued until 30s of expiration time
	if time.Until(claims.ExpiresAt.Time) > 240*time.Second {
		log.Println("NEW REFRESH ONLY BEFORE 4 mins OF EXPIRY")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate a new JWT token
	signedToken, err := generateToken(claims.Username)
	if err != nil {
		log.Println("ERROR OCCURRED WHILE CREATING JWT TOKEN: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Setting JWT claims
	expirationTime := time.Now().Add(5 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "JWtoken",
		Value:   signedToken,
		Expires: expirationTime,
	})

	log.Println("TOKEN REFRESH SUCCESSFUL")
}
