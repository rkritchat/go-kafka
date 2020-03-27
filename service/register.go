package service

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"go-kafka/common"
	"go-kafka/model"
	"go-kafka/str"
	"net/http"
)

type RegisterTask struct {
	model.RegisterRq
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var t RegisterTask
	if err := t.initReq(r); err != nil {
		common.RespErr(w, err)
		return
	}

	u, err := json.Marshal(&t)
	if err != nil {
		common.RespErr(w, err)
		return
	}

	//logging write message
	go t.log(h.p)

	//Write message
	if err := t.Produce(u, h.p); err != nil {
		common.RespErr(w, err)
		return
	}
	common.RespSuccess(w, str.RegisterSuccessFully)
}

func (t *RegisterTask) initReq(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&t)
	return err
}

func (t *RegisterTask) Produce(u []byte, p *kafka.Producer) error {
	topic := viper.GetString("kafka.topic")
	return p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          u,
	}, nil)
}

func (t *RegisterTask) log(p *kafka.Producer) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed , message is : %v\n", string(ev.Value))
			} else {
				fmt.Printf("Delivered success, message is : %v\n", string(ev.Value))
			}
		}
	}
}
