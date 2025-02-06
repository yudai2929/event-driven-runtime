package pubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Subscriber is an interface for Pub/Sub subscribers.
type Subscriber interface {
	Subscribe(ctx context.Context, handler SubscribeHandler)
}

// SubscribeHandler is a function type for handling messages.
type SubscribeHandler func(message string)

type subscriber struct {
	chName string
	rc     *redis.Client
}

type message struct {
}

// NewSubscriber creates a new Pub/Sub subscriber.
func NewSubscriber(addr string, chName string) Subscriber {
	rc := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &subscriber{
		rc:     rc,
		chName: chName,
	}
}

// Subscribe subscribes to a topic.
func (s *subscriber) Subscribe(ctx context.Context, handler SubscribeHandler) {
	pubsub := s.rc.Subscribe(ctx, s.chName)
	defer pubsub.Close()

	// メッセージの受信ループ
	for msg := range pubsub.Channel() {
		handler(msg.Payload)
	}
}
