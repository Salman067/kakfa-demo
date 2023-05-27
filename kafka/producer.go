package kafka

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo/v4"
)

type Comment struct {
	Text string `json:"text"`
}

func connectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func pushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:9092"}
	producer, err := connectProducer(brokersUrl)
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
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func CreateComment(c echo.Context) error {
	cmt := new(Comment)
	if err := c.Bind(cmt); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(cmt)
	cmtInBytes, _ := json.Marshal(cmt)
	err := pushCommentToQueue("comments", cmtInBytes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, cmt)
}
