package main

import (
	"fmt"
	"github.com/cbll/stockmarket-service/lib"
	"net/http"
	"log"
	"firebase.google.com/go"
	"golang.org/x/net/context"
	"strings"
	"encoding/json"
	"google.golang.org/api/option"
)


type myapp struct {
	fbapp *firebase.App
}

func main () {
	ma := &myapp{
		fbapp: InitializeAppWithServiceAccount(),
	}

	go lib.GetStockData()

	http.HandleFunc("/_ah/stocks", ma.StocksHandler)
	http.HandleFunc("/_ah/health", ma.healthCheckHandler)

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func (ma *myapp) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}


func (ma *myapp) StocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	reqToken = splitToken[1]
	lib.VerifyIDToken(ma.fbapp, reqToken)
	enc := json.NewEncoder(w)
	enc.Encode(lib.MarketDataMap)
}


func InitializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account]
	opt := option.WithCredentialsFile("keystore/your-app-firebase-adminsdk.json")
	Application, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return Application
}




