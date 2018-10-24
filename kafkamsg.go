package main

type KafkaMsg struct {
	RemoteAddr string      `json:"remote_addr"`
	RequestURI string      `json:"request_uri"`
	Method     string      `json:"method"`
	Headers    interface{} `json:"headers"`
	Form       interface{} `json:"form"`
	Body       string      `json:"body"`
}
