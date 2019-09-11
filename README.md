# MQ GO HTTP SDK  
Alyun MQ Documents: http://www.aliyun.com/product/ons

Aliyun MQ Console: https://ons.console.aliyun.com


## Use

1. import `github.com/aliyunmq/mq-http-go-sdk`
2. setup GOPATH:
    - MAC/LINUX: export GOPATH={dir}, export GOBIN=$GOPATH/bin
    - WINDOWS: set GOPATH={dir}
2. go get -x -v

## Sample

### V1.0.0 Samples
[Publish Message](https://github.com/aliyunmq/mq-http-samples/blob/master/go/producer.go)

[Consume Message](https://github.com/aliyunmq/mq-http-samples/blob/master/go/consumer.go)

### V1.0.1 Samples
[Publish Message](https://github.com/aliyunmq/mq-http-samples/tree/101-dev/go/producer.go)

[Consume Message](https://github.com/aliyunmq/mq-http-samples/tree/101-dev/go/consumer.go)

[Transaction Message](https://github.com/aliyunmq/mq-http-samples/tree/101-dev/go/trans_producer.go)

Note for 1.0.1: Http consumer only support timer msg(less than 3 days), no matter the msg is produced from http or tcp protocal.

### V1.0.3 Samples
[Publish Message](https://github.com/aliyunmq/mq-http-samples/tree/103-dev/go/producer.go)

[Consume Message](https://github.com/aliyunmq/mq-http-samples/tree/103-dev/go/consumer.go)

[Transaction Message](https://github.com/aliyunmq/mq-http-samples/tree/103-dev/go/trans_producer.go)

[Publish Order Message](https://github.com/aliyunmq/mq-http-samples/tree/103-dev/go/order_producer.go)

[Consume Order Message](https://github.com/aliyunmq/mq-http-samples/tree/103-dev/go/order_consumer.go)

Not for 1.0.3: Order is only supported at special server cluster.
