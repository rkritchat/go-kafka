package main

import (
	"github.com/spf13/viper"
	"go-kafka/common"
	"go-kafka/service"
	"net/http"
)

func main() {
	//Initial DB
	common.InitConfig()
	db := common.InitDB()
	defer db.Close()

	//Initial Kafka Consumer
	c := common.KafkaConsumer()
	defer c.Close()

	s := service.NewHandler(&db, c)
	r := common.InitRouter(s)
	//init consumer
	go s.Consumer()

	panic(http.ListenAndServe(viper.GetString("server.port"), r))
}
