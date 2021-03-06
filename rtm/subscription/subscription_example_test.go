package subscription

import (
	"errors"
	"github.com/satori-com/satori-rtm-sdk-go/logger"
	"github.com/satori-com/satori-rtm-sdk-go/rtm/pdu"
)

// Creates new listener instance and specifies several callbacks
func ExampleNewListener() {
	listener := Listener{
		OnData: func(data pdu.SubscriptionData) {
			// Got messages
			for _, message := range data.Messages {
				logger.Info(string(message))
			}
		},
		OnSubscribeError: func(err pdu.SubscribeError) {
			// Subscribe error
			logger.Error(errors.New(err.Error + "; " + err.Reason))
		},
		OnSubscribed: func(sok pdu.SubscribeOk) {
			logger.Info("Successfully subscribed from position: " + sok.Position)
		},
	}
	logger.Info(listener)
}
