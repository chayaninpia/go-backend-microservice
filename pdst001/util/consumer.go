package util

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func Consumer(topic string, partition int, ch chan<- *kafka.Message) (*kafka.Message, error) {
	// to consume messages
	server := viper.GetString(`kafka.serverAddress`)
	port := viper.GetString(`kafka.serverPort`)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%s", server, port)},
		Topic:     topic,
		Partition: partition,
		MaxBytes:  10e6,
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return nil, fmt.Errorf("read msg failed: %s", err.Error())
		}
		log.Printf("waiting msg ")
		if len(msg.Value) > 0 {
			log.Println("msg coming", string(msg.Value))
			ch <- &msg
			// return &msg, nil
		}
	}
	// conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%s", server, port), topic, partition)
	// if err != nil {
	// 	return fmt.Errorf("failed to dial leader: %s", err.Error())
	// }

	// // conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// b := make([]byte, 10e5) // 100KB max per message
	// for {

	// 	_, err = conn.Read(b)
	// 	if err != nil {
	// 		return fmt.Errorf("read msg failed: %s", err.Error())
	// 	}
	// 	log.Printf("waiting msg ")
	// 	if len(b) > 0 {
	// 		log.Printf(string(b))
	// 		ch <- b
	// 	}
	// }

	// if err := conn.Close(); err != nil {
	// 	return fmt.Errorf("failed to close connection: %s", err)
	// }

	// return nil
}

// func Consumer(c *kafka.Consumer, topic string) *kafka.Message {

// 	// A signal handler or similar could be used to set this to false to break the loop.
// 	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

// 	for {
// 		msg, err := c.ReadMessage(time.Second)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 			return msg
// 		} else if err.(kafka.Error).IsFatal() {
// 			// The client will automatically try to recover from all errors.
// 			// Timeout is not considered an error because it is raised by
// 			// ReadMessage in absence of messages.
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 			return nil
// 		}
// 	}
// }
