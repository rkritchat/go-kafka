package common

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-kafka/path"
	"go-kafka/service"
	"go-kafka/str"
)

func InitDB() gorm.DB{
	db, err := gorm.Open("mysql", viper.GetString("db"))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	return *db
}

func InitConfig(){
	viper.SetConfigName(str.Config)
	viper.AddConfigPath(str.Dot)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func InitRouter(h *service.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(path.Inquiry,h.Inquiry)
	return r
}

func KafkaConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.host"),
		"group.id":          viper.GetString("kafka.group"),
		"auto.offset.reset": viper.GetString("kafka.auto-offset"),
	})

	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics([]string{viper.GetString("kafka.topic"), "^aRegex.*[Tt]opic"}, nil); err != nil{
		panic(err)
	}

	return c
}