/*
Copyright 2017 Mosaic Networks Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package babble

import (
	"fmt"
	"time"
)

type SocketBabbleProxy struct {
	nodeAddress string
	bindAddress string

	client *SocketBabbleProxyClient
	server *SocketBabbleProxyServer
}

func NewSocketBabbleProxy(nodeAddr string, bindAddr string, timeout time.Duration) (*SocketBabbleProxy, error) {
	client := NewSocketBabbleProxyClient(nodeAddr, timeout)
	server, err := NewSocketBabbleProxyServer(bindAddr)
	if err != nil {
		return nil, err
	}

	proxy := &SocketBabbleProxy{
		nodeAddress: nodeAddr,
		bindAddress: bindAddr,
		client:      client,
		server:      server,
	}
	go proxy.server.listen()

	return proxy, nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//Implement BabbleProxy interface

func (p *SocketBabbleProxy) CommitCh() chan []byte {
	return p.server.commitCh
}

func (p *SocketBabbleProxy) SubmitTx(tx []byte) error {
	ack, err := p.client.SubmitTx(tx)
	if err != nil {
		return err
	}
	if !*ack {
		return fmt.Errorf("Failed to deliver transaction to Babble")
	}
	return nil
}
