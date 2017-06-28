package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"io/ioutil"

	"encoding/json"

	"bufio"

	"github.com/streadway/amqp"
)

type appConfig struct {
	Amqp  string `json:"amqp"`
	Queue string `json:"queue"`
}

var app appConfig

func init() {
	data, e := ioutil.ReadFile("./app.json")
	failOnError(e, "unread file app.json")
	e = json.Unmarshal(data, &app)
	failOnError(e, "unmarshal data to struct")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial(app.Amqp)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	dataFile, err := os.Open("./data.txt")
	failOnError(err, "Failed to open a file: data.txt")
	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)

	fmt.Printf("目标队列名为: %v\n", app.Queue)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && len(strings.TrimSpace(line)) > 0 {
			err = ch.Publish(
				"",        // exchange
				app.Queue, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(line),
				})
			failOnError(err, "send data ERRER =>"+line)
			fmt.Printf("send data OK =>%v\n", line)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Done.")
}
