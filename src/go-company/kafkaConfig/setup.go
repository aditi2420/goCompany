package kafkaconfig

import (
	"context"
	"go-company/configs"
	"go-company/models"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var KafkaConn *kafka.Conn

type MessagePayload struct {
	EventType string
	Payload models.Company
	Time time.Time	
}

const(
	CompanyCreated string = "company_created"
	CompanyDeleted string = "company_deleted"
	CompanyUpdated string = "company_updated"
	//CompanyCreated string = "company_created"
)

func SetupKafkaProducer() {

	topic := configs.KafkaTopic
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to connect to kafka producer topic:", err)
		log.Panic(err.Error())
	}
	
	//conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	KafkaConn = conn

}

func WriteToProducer(payload []byte) {
	KafkaConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := KafkaConn.WriteMessages(
		kafka.Message{Topic: configs.KafkaTopic,
			Key:   []byte("company"),
			Value: payload,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
