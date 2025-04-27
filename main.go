package main

import (
	controller "codeproc/controller/http"
	"codeproc/infrastructure/consumer/codeprocessor"
	"codeproc/infrastructure/repository/ram_storage"
	pkgHttp "codeproc/pkg/http"
	"codeproc/usecases/service"
	"flag"
	"log"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title CodeProcessor
// @version 1.0
// @description This is a HTTP-server for code processing.

// @host localhost:8080
// @BasePath /task

func main() {
	addr := flag.String("addr", ":8080", "address for http server")

	storage := ram_storage.NewStorage()
	consumer := codeprocessor.NewConsumer()
	service := service.NewObject(storage, consumer)
	server := controller.New(service)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	server.WithObjectHandler(r)

	if err := pkgHttp.CreateAndRunServer(*addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
