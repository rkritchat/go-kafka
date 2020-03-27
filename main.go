package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go-kafka/config"
	"go-kafka/path"
	"go-kafka/service"
	"net/http"
)

func main() {
	//Initial DB
	config.InitConfig()
	db := config.InitDB()
	defer db.Close()

	//Initial Kafka Consumer
	c := config.KafkaConsumer()
	defer c.Close()
	//Initial Kafka Producer
	p := config.KafkaProducer()
	defer p.Close()

	s := service.NewHandlers(&db, c, p)
	r := initRouter(s)
	//init consumer
	go s.Consumer()

	panic(http.ListenAndServe(viper.GetString("server.port"), r))
}

func initRouter(h *service.Handlers) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(path.Inquiry, h.Inquiry).Methods(http.MethodGet)
	r.HandleFunc(path.Register, h.Register).Methods(http.MethodPost)
	return r
}
