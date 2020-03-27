package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type Handlers struct {
	db *gorm.DB
	c  *kafka.Consumer
	p  *kafka.Producer
}

func NewHandlers(db *gorm.DB, c *kafka.Consumer, p *kafka.Producer) *Handlers {
	return &Handlers{
		db: db,
		c:  c,
		p:  p,
	}
}
