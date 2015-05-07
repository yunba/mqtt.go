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
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"log"
	MQTT "mqtt"
)


var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	stdin := bufio.NewReader(os.Stdin)

	appkey := "13213131"
	deviceId := ""
	topic := "underground"
	qos := MQTT.QOS_ONE

	yunbaClient := &MQTT.YunbaClient{appkey, deviceId}
	regInfo, err := yunbaClient.Reg()
	if err != nil {
		log.Fatal(err)
		return
	}

	if regInfo.ErrCode != 0 {
		log.Fatal("has error:", regInfo.ErrCode)
		return
	}

	fmt.Printf("resp:\t\t%+v\n", regInfo)
	fmt.Println("ClientId", regInfo.Client)
	fmt.Println("UserName", regInfo.UserName)
	fmt.Println("Password", regInfo.Password)
	fmt.Println("DeviceId", regInfo.DeviceId)
	fmt.Println("")

	urlInfo, err := yunbaClient.GetHost()
	if err != nil {
		log.Fatal(err)
		return
	}
	if regInfo.ErrCode != 0 {
		log.Fatal("has error:", urlInfo.ErrCode)
		return
	}


	fmt.Printf("URL:\t\t%+v\n", urlInfo)
	fmt.Println("url", urlInfo.Client)
	fmt.Println("")


	connOpts := MQTT.NewClientOptions()
	connOpts.AddBroker(urlInfo.Client)
	connOpts.SetClientId(regInfo.Client)
	connOpts.SetCleanSession(true)
	connOpts.SetProtocolVersion(0x13)

	connOpts.SetUsername(regInfo.UserName)
	connOpts.SetPassword(regInfo.Password)

	connOpts.SetDefaultPublishHandler(f)

	client := MQTT.NewClient(connOpts)
	_, err = client.Start()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Connected to %s\n", urlInfo.Client)
	}

	for {
		message, err := stdin.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}
		r := client.Publish(qos, topic, []byte(strings.TrimSpace(message)))
		<-r // received puback will send message to chan r,   net.go: case PUBACK
		fmt.Println("Message Sent")
	}
}