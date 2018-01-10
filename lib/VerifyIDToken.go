package lib

import (
	"log"
	"golang.org/x/net/context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
)


func VerifyIDToken(app *firebase.App, idToken string) *auth.Token {
	// [START verify_id_token]
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	// [END verify_id_token]

	return token
}

