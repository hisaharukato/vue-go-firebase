package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
    "google.golang.org/api/option"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Firebase SDK のセットアップ
        opt := option.WithCredentialsFile("./vue-golang-a3fb7-firebase-adminsdk-cb8dj-770eab8c08.json")
        app, err := firebase.NewApp(context.Background(), nil, opt)
        if err != nil {
					fmt.Printf("errorNewApp: %v\n", err)
					os.Exit(1)
				}
        auth, err := app.Auth(context.Background())
        if err != nil {
            fmt.Printf("errorAuth: %v\n", err)
            os.Exit(1)
        }

        // クライアントから送られてきた JWT 取得
        authHeader := r.Header.Get("Authorization")
        idToken := strings.Replace(authHeader, "Bearer ", "", 1)

        // JWT の検証
        token, err := auth.VerifyIDToken(context.Background(), idToken)
        if err != nil {
            // JWT が無効なら Handler に進まず別処理
            fmt.Printf("error verifying ID token: %v\n", err)
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("error verifying ID token\n"))
            return
        }
        log.Printf("Verified ID token: %v\n", token)
        next.ServeHTTP(w, r)
    }
}

func public(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello public!あううあいおおいううあ\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello private!\n"))
}

func main(){
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	r.HandleFunc("/private", authMiddleware(private))

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}