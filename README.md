# Test
$ go test

# publish
$ go run stdinpub.go -appkey xxx -topic yyy -qos 1

# subscribe
$ go run stdoutsub.go -appkey xxx -topic yyy -qos 1

# set alias and publish to alias
$ go run alias.go -appkey xxx -alias zzz


mqtt.go
=======

mqtt library in golang

Eclipse Paho MQTT Go client
===========================


This repository contains the source code for the [Eclipse Paho](http://eclipse.org/paho) MQTT Go client library. 

This code builds a library which enable applications to connect to an [MQTT](http://mqtt.org) broker to publish messages, and to subscribe to topics and receive published messages.

This library supports a fully asynchronous mode of operation.


Installation and Build
----------------------

This client is designed to work with the standard Go tools, so installation is as easy as:

```
go get github.com/yunba/mqtt.go
```

Usage and API
-------------

Detailed API documentation is available by using to godoc tool, or can be browsed online
using the [godoc.org](http://godoc.org/git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git) service.

Make use of the library by importing it in your Go client source code. For example,

```
MQTT "github.com/yunba/mqtt.go"
```

Samples are available in the `/samples` directory for reference.


Runtime tracing
---------------

Tracing is enabled by using the `SetTraceLevel` option when creating a ClientOptions struct. See the ClientOptions
documentation for more details.


Reporting bugs
--------------

Please report bugs under the "MQTT-Go" Component in [Eclipse Bugzilla](http://bugs.eclipse.org/bugs/) for the Paho Technology project. This is a very new library as of Q1 2014, so there are sure to be bugs.


More information
----------------

Discussion of the Paho clients takes place on the [Eclipse paho-dev mailing list](https://dev.eclipse.org/mailman/listinfo/paho-dev).

General questions about the MQTT protocol are discussed in the [MQTT Google Group](https://groups.google.com/forum/?hl=en-US&fromgroups#!forum/mqtt).

There is much more information available via the [MQTT community site](http://mqtt.org).

