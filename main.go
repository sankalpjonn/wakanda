package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/sankalpjonn/wrq"
)

var (
	serverhost string
	brokerlist string
)

const (
	MAX_WORKERS = 10
)

func startServer(serverAddr string, dispatcher *wrq.Dispatcher, producer *kafka.Producer) {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/topics/:topic", handler(dispatcher, producer))
	r.POST("/topics/:topic", handler(dispatcher, producer))
	r.PUT("/topics/:topic", handler(dispatcher, producer))
	r.PATCH("/topics/:topic", handler(dispatcher, producer))
	r.Run(serverAddr)
}

func startWakanda(serverhost string, brokerlist string) {
	dispatcher := wrq.New()
	defer dispatcher.Stop()

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokerlist})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	go startServer(serverhost, dispatcher, producer)

	// wait for kill signal
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sigc
}

func main() {
	flag.StringVar(&serverhost, "host", "0.0.0.0:8000", "Host for the wakanda server to listen on")
	flag.StringVar(&brokerlist, "broker-list", "127.0.0.1:9092", "Kafka broker list")
	flag.Parse()

	startWakanda(serverhost, brokerlist)
}
