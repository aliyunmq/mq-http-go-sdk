package mq_http_sdk

import (
	"fmt"
	"strings"
)

// MQ的消息生产者
type MQProducer interface {
	// 主题名字
	TopicName() string
	// 实例ID，可空
	InstanceId() string
	// 发送消息
	PublishMessage(message PublishMessageRequest) (resp PublishMessageResponse, err error)
}

type AliyunMQProducer struct {
	topicName  string
	instanceId string
	client     *AliyunMQClient
	decoder    MQDecoder
}

func (p *AliyunMQProducer) TopicName() string {
	return p.topicName
}

func (p *AliyunMQProducer) InstanceId() string {
	return p.instanceId
}

func (p *AliyunMQProducer) PublishMessage(message PublishMessageRequest) (resp PublishMessageResponse, err error) {
	resourceBuilder := strings.Builder{}
	if p.instanceId != "" {
		resourceBuilder.WriteString(fmt.Sprintf("topics/%s/%s?ns=%s", p.topicName, "messages", p.instanceId))
	} else {
		resourceBuilder.WriteString(fmt.Sprintf("topics/%s/%s", p.topicName, "messages"))
	}
	_, err = p.client.Send(p.decoder, POST, nil, message, resourceBuilder.String(), &resp)
	return
}
