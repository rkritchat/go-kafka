package service

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"go-kafka/entity"
	"go-kafka/model"
	"go-kafka/rule"
)

type ConsumerTask struct {
	db *gorm.DB
	u *model.UserInfo
}

func (h *Handler)Consumer(){
	var task ConsumerTask
	task.db = h.db
	for{
		msg, err := h.c.ReadMessage(-1)
		if err== nil{
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			task.execute(string(msg.Value))
		}else{
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (t *ConsumerTask)execute(s string) {
	 if err := t.initJson(s); err != nil {
		 fmt.Printf("err occrred  (%v)\n", err)
		 return
	 }
	 if err := t.valid(); err!=nil{
		 fmt.Printf("err occrred (%v)\n", err)
		 return
	 }
	 t.save()
}

func (t *ConsumerTask)initJson(s string) error {
	var j model.UserInfo
	err := json.Unmarshal([]byte(s), &j)
	t.u = &j
	return err
}

func (t *ConsumerTask)valid() error {
	 return validation.ValidateStruct(t.u,
	 	validation.Field(&t.u.FistName, rule.FistName...),
	 	validation.Field(&t.u.LastName, rule.LastName...),
	 	validation.Field(&t.u.Age, rule.Age...))
}

func (t *ConsumerTask)save(){
	save := t.db.Save(&entity.UserInfo{
		FirstName: t.u.FistName,
		LastName:  t.u.LastName,
		Age:       t.u.Age,
	})
	if save.Error != nil{
		fmt.Printf("err occrred  (%v)\n", save.Error)
	}
}


