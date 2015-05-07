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

	appkey := "13215315"
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

	client := MQTT.NewClient(connOpts)
	_, err = client.Start()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Connected to %s\n", urlInfo.Client)
	}

	filter, e := MQTT.NewTopicFilter(topic, byte(qos))
	if e != nil {
		log.Fatal(e)
	}
	client.StartSubscription(onMessageReceived, filter)

	for {
		time.Sleep(1 * time.Second)
	}
}