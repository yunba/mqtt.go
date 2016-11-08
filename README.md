# mqtt.go

Yunba Go SDK.

## Installation

This client is designed to work with the standard Go tools, so installation is as easy as:

```
go get github.com/yunba/mqtt.go
```

## Test

```
go test
```


## Usage

Make use of the library by importing it in your Go client source code. For example,

```
MQTT "github.com/yunba/mqtt.go"
```

Samples are available in the `samples` directory for reference.

## Samples

_publish_

```
cd samples
go run stdinpub.go -appkey xxx -topic yyy -qos 1
```

_subscribe_

```
go run stdoutsub.go -appkey xxx -topic yyy -qos 1
```

_set alias and publish to alias_

```
go run alias.go -appkey xxx -alias zzz
```

## Eclipse Paho MQTT Go client

This repository contains the source code for the [Eclipse Paho](http://eclipse.org/paho) MQTT Go client library.


