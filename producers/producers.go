package producers

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	// config.Producer.Return.Successes = true
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)

	if err != nil {
		log.Fatalln("Open Kafka Connection Failed")
		return nil, err
	}

	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {

	brokersUrl := []string{"http://localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("send message failed", err)
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
