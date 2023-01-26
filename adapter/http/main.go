package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/EliasSantiago/app-bank-transactions/adapter/postgres"
	"github.com/EliasSantiago/app-bank-transactions/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()
	postgres.RunMigrations()
	userService := di.ConfigUserDI(conn)
	walletService := di.ConfigWalletDI(conn)
	transactionService := di.ConfigTransactionDI(conn)
	router := mux.NewRouter()
	router.Handle("/user", http.HandlerFunc(userService.Create)).Methods("POST")
	router.Handle("/wallet", http.HandlerFunc(walletService.Create)).Methods("POST")
	router.Handle("/transaction", http.HandlerFunc(transactionService.Create)).Methods("POST")
	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
