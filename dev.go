package main

import(
	"fmt"
	"os"
	"github.com/Shopify/sarama"
)

func main() {

	producer, consumer := conf()
    
    publish(producer)

    partitions, _ := consumer.Partitions("verifiedData")
    consume, err := consumer.ConsumePartition("verifiedData", partitions[0], sarama.OffsetNewest)
	
	if nil != err {
		fmt.Printf("Topic %v Partitions: %v", "verifiedData", partitions)
		panic(err)
	}

    msg := <-consume.Messages()
    fmt.Println("Value: ", string(msg.Key), string(msg.Value))
}


func conf()(sarama.SyncProducer, sarama.Consumer) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true

	kafkaBroker := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(kafkaBroker, config)
	consumer, err := sarama.NewConsumer(kafkaBroker, config)

	if err != nil {
        fmt.Println("Error producer: ", err.Error())
        os.Exit(1)
    }
    return producer, consumer	
}


func publish(producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
	Topic: "userData",
	Value: sarama.StringEncoder("+94772564795"),
	}
	
	producer.SendMessage(msg)
}