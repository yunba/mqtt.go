/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

import MQTT "github.com/yunba/mqtt.go"

var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	payload := msg.Payload()
	data := string(payload[4:])
	fmt.Println("data is: ", data)
}

func main() {
	//stdin := bufio.NewReader(os.Stdin)
	hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://127.0.0.1:1883", "The full URL of the MQTT server to connect to")
	//retained := flag.Bool("retained", false, "Are the messages sent with the retained flag")
	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", "", "A username to authenticate to the MQTT server")
	password := flag.String("password", "", "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientId(*clientid).SetCleanSession(true).SetProtocolVersion(0x13)
	connOpts.SetDefaultPublishHandler(f)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}

	client := MQTT.NewClient(connOpts)
	_, err := client.Start()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	client.GetAlias()

	for {
		time.Sleep(1 * time.Second)
	}
}
