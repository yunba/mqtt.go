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

package mqtt

import (
	"log"
	"os"
	"fmt"
)

var (
	ERROR    *log.Logger
	CRITICAL *log.Logger
	WARN     *log.Logger
	DEBUG    *log.Logger
)

func init() {
	logfile,err := os.OpenFile("mqtt.log",os.O_RDWR|os.O_CREATE, os.ModeAppend | 0666);
    if err!=nil {
        fmt.Printf("%s\r\n",err.Error());
        os.Exit(-1);
    }
	ERROR = log.New(logfile, "", 0)
	CRITICAL = log.New(logfile, "", 0)
	WARN = log.New(logfile, "", 0)
	DEBUG = log.New(logfile, "", 0)
}
