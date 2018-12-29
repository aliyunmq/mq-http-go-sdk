package mq_http_sdk

import (
	"encoding/xml"
)

type MessageResponse struct {
	XMLName xml.Name `xml:"Message" json:"-"`
	// 错误码
	Code string `xml:"Code,omitempty" json:"code,omitempty"`
	// 错误描述
	Message string `xml:"Message,omitempty" json:"message,omitempty"`
	// 请求ID
	RequestId string `xml:"RequestId,omitempty" json:"request_id,omitempty"`
	// 请求HOST
	HostId string `xml:"HostId,omitempty" json:"host_id,omitempty"`
}

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error" json:"-"`
	// 错误码
	Code string `xml:"Code,omitempty" json:"code,omitempty"`
	// 错误描述
	Message string `xml:"Message,omitempty" json:"message,omitempty"`
	// 请求ID
	RequestId string `xml:"RequestId,omitempty" json:"request_id,omitempty"`
	// 请求HOST
	HostId string `xml:"HostId,omitempty" json:"host_id,omitempty"`
}

type PublishMessageRequest struct {
	XMLName xml.Name `xml:"Message" json:"-"`
	// 消息体
	MessageBody string `xml:"MessageBody" json:"message_body"`
	// 消息标签
	MessageTag string `xml:"MessageTag,omitempty" json:"message_tag,omitempty"`
}

type PublishMessageResponse struct {
	MessageResponse
	// 消息ID
	MessageId string `xml:"MessageId" json:"message_id"`
	// 消息体MD5
	MessageBodyMD5 string `xml:"MessageBodyMD5" json:"message_body_md5"`
}

type ReceiptHandles struct {
	XMLName xml.Name `xml:"ReceiptHandles"`
	// 消息句柄
	ReceiptHandles []string `xml:"ReceiptHandle"`
}

type AckMessageErrorEntry struct {
	XMLName xml.Name `xml:"Error" json:"-"`
	// Ack消息出错的错误码
	ErrorCode string `xml:"ErrorCode" json:"error_code"`
	// Ack消息出错的错误描述
	ErrorMessage string `xml:"ErrorMessage" json:"error_messages"`
	// Ack消息出错的消息句柄
	ReceiptHandle string `xml:"ReceiptHandle,omitempty" json:"receipt_handle"`
}

type AckMessageErrorResponse struct {
	XMLName        xml.Name               `xml:"Errors" json:"-"`
	FailedMessages []AckMessageErrorEntry `xml:"Error" json:"errors"`
}

type ConsumeMessageEntry struct {
	MessageResponse
	// 消息ID
	MessageId string `xml:"MessageId" json:"message_id"`
	// 消息句柄
	ReceiptHandle string `xml:"ReceiptHandle" json:"receipt_handle"`
	// 消息体MD5
	MessageBodyMD5 string `xml:"MessageBodyMD5" json:"message_body_md5"`
	// 消息体
	MessageBody string `xml:"MessageBody" json:"message_body"`
	// 消息发送时间
	PublishTime int64 `xml:"PublishTime" json:"publish_time"`
	// 下次消费消息的时间（如果这次消费的消息没有Ack）
	NextConsumeTime int64 `xml:"NextConsumeTime" json:"next_consume_time"`
	// 第一次消费的时间
	FirstConsumeTime int64 `xml:"FirstConsumeTime" json:"first_consume_time"`
	// 消费次数
	ConsumedTimes int64 `xml:"ConsumedTimes" json:"consumed_times"`
	// 消息标签
	MessageTag string `xml:"MessageTag" json:"message_tag"`
}

type ConsumeMessageResponse struct {
	XMLName  xml.Name              `xml:"Messages" json:"-"`
	Messages []ConsumeMessageEntry `xml:"Message" json:"messages"`
}
