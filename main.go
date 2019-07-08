package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/wedgwoodtom/go-common-module-module/awsClients"
	"github.com/wedgwoodtom/go-dep-client/messageProcessor"
)

// Values here are set during the build
var VERSION = "N/A"

const (
	aws_region = "us-west-2"
	sqs_queue  = "urcs"
)

//noinspection GoUnhandledErrorResult
func main() {
	// Super simple hello world to get started
	http.HandleFunc("/urcs", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello WORLD, I am URCS")
	})

	// alive check
	http.HandleFunc("/urcs/management/alive", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Web Service is Ok")
	})

	log.Println("Starting Message Queue Processor")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(aws_region),
	}))
	sqsMessageQueue := awsClients.NewQueue(sqs.New(sess), sqs_queue)
	messageProcessor := messageProcessor.New(sqsMessageQueue)
	go messageProcessor.Start()

	printBanner()
	log.Fatal(http.ListenAndServe(":10532", nil))
}

func printBanner() {
	log.Println(
		`
--------------------------
  _   _ ____   ____ ____  
 | | | |  _ \ / ___/ ___| 
 | | | | |_) | |   \___ \    is RUNNING...
 | |_| |  _ <| |___ ___) |
  \___/|_| \_\\____|____/ 
--------------------------`)
	log.Println("Version: ", VERSION)
}
