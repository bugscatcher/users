package testutil

import (
	"github.com/Shopify/sarama"
)

type KafkaMock struct {
	expectedCall int
	messages     map[string][]*sarama.ProducerMessage
}

func MockKafkaProducer(expectedCallCount int) *KafkaMock {
	mock := KafkaMock{}
	mock.expectedCall = expectedCallCount
	mock.messages = make(map[string][]*sarama.ProducerMessage)
	return &mock
}

func (k *KafkaMock) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	if k.expectedCall <= 0 {
		panic("unexpected method call")
	}
	k.expectedCall--
	topic := msg.Topic
	messages := k.messages[topic]
	if messages == nil {
		messages = make([]*sarama.ProducerMessage, 0, 1)
	}
	messages = append(messages, msg)
	k.messages[topic] = messages
	return 0, 1, nil

}

func (k *KafkaMock) SendMessages(msgs []*sarama.ProducerMessage) error {
	for _, msg := range msgs {
		_, _, _ = k.SendMessage(msg)
	}
	return nil
}

func (k *KafkaMock) Close() error {
	return nil
}
