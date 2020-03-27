package config

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-kafka/str"
)

func InitConfig() {
	viper.SetConfigName(str.Config)
	viper.AddConfigPath(str.Dot)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func InitDB() gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("db"))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	return *db
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

	if err := c.SubscribeTopics([]string{viper.GetString("kafka.topic"), "^aRegex.*[Tt]opic"}, nil); err != nil {
		panic(err)
	}

	return c
}

func KafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.host"),
	})

	if err != nil {
		panic(err)
	}

	return p
}
