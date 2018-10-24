package main

import (
	"io/ioutil"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/sankalpjonn/wrq"
)

func handler(dispatcher *wrq.Dispatcher, producer *kafka.Producer) gin.HandlerFunc {

	fn := func(ginContext *gin.Context) {
		r := ginContext.Request
		r.ParseForm()
		body := r.Body
		x, _ := ioutil.ReadAll(body)
		kafkaMsg := KafkaMsg{
			Headers:    r.Header,
			RemoteAddr: r.RemoteAddr,
			RequestURI: r.RequestURI,
			Method:     r.Method,
			Form:       r.Form,
			Body:       string(x),
		}
		j := &job{
			topic:    ginContext.Param("topic"),
			msg:      kafkaMsg,
			producer: producer,
		}
		dispatcher.AddJob(j)

		ginContext.Header("content-type", "application/json;charset=utf-8")
		ginContext.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
	return gin.HandlerFunc(fn)
}
