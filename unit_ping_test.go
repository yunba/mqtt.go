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

import "testing"

func testNewPingReqMessage(t *testing.T, version byte){
	pr := newPingReqMsg()
	if pr.msgType() != PINGREQ {
		t.Errorf("NewPingReqMessage bad msg type: %v", pr.msgType())
	}
	if pr.remLen() != 0 {
		t.Errorf("NewPingReqMessage bad remlen, expected 0, got %d", pr.remLen())
	}

	exp := []byte{
		0xC0,
		0x00,
	}

	bs := pr.Bytes(version)

	if len(bs) != 2 {
		t.Errorf("NewPingReqMessage.Bytes() wrong length: %d", len(bs))
	}

	if exp[0] != bs[0] || exp[1] != bs[1] {
		t.Errorf("NewPingMessage.Bytes() wrong")
	}
}

func Test_NewPingReqMessage_v3(t *testing.T) {
	testNewPingReqMessage(t, 0x03)
}

func Test_NewPingReqMessage_v19(t *testing.T) {
	testNewPingReqMessage(t, 0x13)
}


func testDecodeMessagePingResp(t *testing.T, version byte){
	bs := []byte{
		0xD0,
		0x00,
	}
	presp := decode(bs, version)
	if presp.msgType() != PINGRESP {
		t.Errorf("DecodeMessage ping response wrong msg type: %v", presp.msgType())
	}
	if presp.remLen() != 0 {
		t.Errorf("DecodeMessage ping response wrong rem len: %d", presp.remLen())
	}

}
func Test_DecodeMessage_Pingresp_v3(t *testing.T) {
	testDecodeMessagePingResp(t, 0x03)
}

func Test_DecodeMessage_Pingresp_v19(t *testing.T) {
	testDecodeMessagePingResp(t, 0x13)
}
