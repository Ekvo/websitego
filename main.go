package main

import (
	"context"
	"fmt"
	"github.com/Ekvo/websitego/internal/application"
	"github.com/Ekvo/websitego/internal/repository"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	dbpool, err := repository.InitDBConn(ctx)
	if err != nil {
		log.Fatalf("%w failed to init DB connection", err)
	}
	defer dbpool.Close()
	a := application.NewApp(ctx, dbpool)
	r := httprouter.New()
	a.Routes(r)
	srv := &http.Server{Addr: "127.0.0.1:8000", Handler: r}
	fmt.Println("It is alive! Try http://localhost:8000")
	srv.ListenAndServe()
}
