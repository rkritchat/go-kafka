package main

import (
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
   //defer c.Close()

   s := service.NewHandler(&db,c)
   s.Consumer()
   r := common.InitRouter(s)
   panic(http.ListenAndServe(":8080", r))
}