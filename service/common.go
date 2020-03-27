package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db *gorm.DB
	c *kafka.Consumer
}

func NewHandler(db *gorm.DB, c *kafka.Consumer) *Handler{
	return &Handler{
		db: db,
		c: c,
	}
}