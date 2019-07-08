package messageProcessor

import (
	"github.com/wedgwoodtom/go-common-module-module/awsClients"
	"log"
)

type MessageProcessor struct {
	messageQueue awsClients.SQSMessageQueue
}

func New(messageQueue awsClients.SQSMessageQueue) MessageProcessor {
	return MessageProcessor{messageQueue: messageQueue}
}

func (mp *MessageProcessor) Start() {
	for {
		mp.processMessages()
	}
}

func (mp *MessageProcessor) processMessages() {
	messages, err := mp.messageQueue.ReceiveMessage(20, awsClients.MaxNumberOfMessages(10))
	if err != nil {
		log.Println("Error in receiving messages ", err)
	}
	if messages == nil {
		log.Println("No messages received")
	}

	for _, message := range messages {
		log.Println("Processing Message: ", message.MessageId)

		// Compute and store to Dynamo

		// TODO: Don't delete for now, but this works
		//log.Println("Deleting Message: ", message.MessageId)
		//deleteError := mp.messageQueue.DeleteMessage(message.ReceiptHandle)
		//if deleteError != nil {
		//	log.Println("Error in Deleting processed message for ", message.ReceiptHandle, deleteError)
		//}
	}
}
