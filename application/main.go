package main

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v4"
)

const (
    baseUrl = "http://localhost:8080/"
    clientId = "test-client"
    clientSecret = "F34gjpncG8iqSv6hgM3lnnKMQdzlEXlk"
    realm = "test_realm"


)


func main(){
	http.HandleFunc("/health", protectResource(callHealthCheck))

	log.Default().Println("server is running in port 8081")
    //login()
	http.ListenAndServe(":8081", nil)
}

func callHealthCheck(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")

	if request.Method == "GET"{
		response.WriteHeader(http.StatusOK)	
		log.Default().Println("Rertuned status code ok")
		status := HealthCheck{
			Status: "healthy",
		}

		json.NewEncoder(response).Encode(status)


	}else {
		log.Default().Println("client error method not allowed")

		response.WriteHeader(http.StatusMethodNotAllowed)
	}

}

type HealthCheck struct {
	Status string `json:"status"`
}


func login() {
    client := gocloak.NewClient(baseUrl)

    jwt, err := client.LoginClient(context.TODO(),clientId,clientSecret,realm)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(jwt)
}


func protectResource(next http.HandlerFunc ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        if token == "" {
            w.WriteHeader(401)

            http.Error(w,"No token provided", http.StatusUnauthorized)
            return
        }

        client := gocloak.NewClient(baseUrl)
        retrieved_token, _, err := client.DecodeAccessToken(context.TODO(),token,realm)
        if err != nil {
            w.WriteHeader(401)
            http.Error(w,fmt.Sprint(err), http.StatusUnauthorized)
            fmt.Println(err)
            return
        }

        custom_claim := "scope"
        required_scope := "health_check_scope"

        claims := retrieved_token.Claims.(jwt.MapClaims) // having serios doubts in this line

        if claims[custom_claim] ==  required_scope {

            next(w,r)
        } else{
            w.WriteHeader(401)

            http.Error(w,"not authorized",http.StatusUnauthorized)
            return
        }
    }

}

