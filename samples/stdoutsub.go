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
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"

	MQTT "mqtt"
	"flag"
	"strconv"
)

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("signal received, exiting")
		os.Exit(0)
	}()

	hostname, _ := os.Hostname()

	appkey := flag.String("appkey", "", "YunBa appkey")
	topic := flag.String("topic", hostname, "Topic to publish the messages on")
	qos := flag.Int("qos", 0, "The QoS to send the messages at")
	//retained := flag.Bool("retained", false, "Are the messages sent with the retained flag")
	deviceId := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	flag.Parse()

	if *appkey == "" {
		log.Fatal("please set your Yunba Portal's appkey")
	}

	yunbaClient := &MQTT.YunbaClient{*appkey, *deviceId}
	regInfo, err := yunbaClient.Reg()
	if err != nil {
		log.Fatal(err)
	}

	if regInfo.ErrCode != 0 {
		log.Fatal("reg has error:", regInfo.ErrCode)
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
	}
	if regInfo.ErrCode != 0 {
		log.Fatal("reg has error:", urlInfo.ErrCode)
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


	client := MQTT.NewClient(connOpts)
	_, err = client.Start()
	if err != nil {
		panic(err)
	} else {
        log.Printf("Connected to %s\n", urlInfo.Client)
    }

    <- client.SetAlias(hostname)

    filter, e := MQTT.NewTopicFilter(*topic, byte(*qos))
	if e != nil {
		log.Fatal(e)
	}

    client.StartSubscription(onMessageReceived, filter)
    client.Presence(onMessageReceived, *topic)

    for {
		time.Sleep(1 * time.Second)
	}
}