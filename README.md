# Wakanda

A service that produces logs to a kafka cluster on a particular topic when hit on the endpoint `/topics/:topic` with a throughput of 5000 requests/sec

![knowhere](http://66.42.57.109/wakanda.jpg)

## Installation
```sh
go get github.com/sankalpjonn/wakanda
```

## Usage

To run wakanda
```sh
wakanda -host 0.0.0.0:8080 -broker-list 127.0.0.1:9092
```
To produce logs
```curl
curl -XPOST http://localhost:8080/topics/test?param1=1&param2=2 -H "Content-Type: application/json" -d '{"test": "value"}'
```

The above curl produces the following log to kafka
```json
{
	"remote_addr": "127.0.0.1:35954",
	"request_uri": "/topics/test?param1=1\u0026param2=2",
	"method": "POST",
	"headers": {
		"Accept": ["*/*"],
		"Content-Length": ["17"],
		"Content-Type": ["application/json"],
		"User-Agent": ["curl/7.47.0"]
	},
	"form": {
		"param1": ["1"],
		"param2": ["2"]
	},
	"body": "{\"test\": \"value\"}"
}
```

## Contact
[sankalpjonna@gmail.com](sankalpjonna@gmail.com)
