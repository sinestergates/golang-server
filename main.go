package main

import (
	"fmt"
    "log"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/go-redis/redis"
    "time"
)

var jwtKey = []byte("123")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func homePage(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to the home page"))
}

func ConnectDB() bool{

    client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес сервера Redis
		Password: "",                // Пароль (если не установлен, оставьте пустым)
		DB:       0,                 // Номер базы данных
	})

	// Проверяем подключение к Redis
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Ошибка при подключении к Redis:", err)
		return false
	}

	fmt.Println("Подключение к Redis успешно:", pong)
    return true
}

func reg_user() string{
    
}

func login(w http.ResponseWriter, r *http.Request) {
    // Authenticating the user, for simplicity taking "username" and "password" in the query parameters
    username := r.URL.Query().Get("username")
    password := r.URL.Query().Get("password")
	fmt.Printf("Error signing the JWT token: %v\n password %v\n", r.Header.Get("Accept-Encoding"), password)
    // Check username and password


    if ConnectDB(){
        expirationTime := time.Now().Add(5 * time.Minute)

        // Create JWT token
        claims := &Claims{
            Username: username,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: expirationTime.Unix(),
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        fmt.Printf("token %v\n", claims, token)
        tokenString, err := token.SignedString(jwtKey)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.Write([]byte(tokenString))
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/login", login)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}