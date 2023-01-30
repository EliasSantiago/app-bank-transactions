package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/EliasSantiago/app-bank-transactions/adapter/http/docs"
	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres"
	"github.com/EliasSantiago/app-bank-transactions/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Elias Fonseca
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	arg1 := os.Args[1]
	println(arg1)
	if arg1 == "api" {
		ctx := context.Background()
		conn := postgres.GetConnection(ctx)
		defer conn.Close()
		postgres.RunMigrations()
		userService := di.ConfigUserDI(conn)
		walletService := di.ConfigWalletDI(conn)
		transactionService := di.ConfigTransactionDI(conn)
		router := mux.NewRouter()
		router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
		router.Handle("/user", http.HandlerFunc(userService.Create)).Methods("POST")
		router.Handle("/wallet", http.HandlerFunc(walletService.Create)).Methods("POST")
		router.Handle("/transaction", http.HandlerFunc(transactionService.Create)).Methods("POST")
		port := viper.GetString("server.port")
		log.Printf("LISTEN ON PORT: %v", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	}
	if arg1 == "consumer" {
		//TODO RUN CONSUMER
		println("Run consumer...")
	}
}
