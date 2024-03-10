package network

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	var addr1 NetAddr = "addr1"
	var addr2 NetAddr = "addr2"

	t1 := NewLocalTransport(addr1)
	t2 := NewLocalTransport(addr2)

	t1.Connect(t2)
	t2.Connect(t1)

	assert.Equal(t, t1.Peers[addr2], t2)
	assert.Equal(t, t2.Peers[addr1], t1)
}

func TestSendMessage(t *testing.T) {
	var addr1 NetAddr = "addr1"
	var addr2 NetAddr = "addr2"

	t1 := NewLocalTransport(addr1)
	t2 := NewLocalTransport(addr2)

	t1.Connect(t2)
	t2.Connect(t1)

	msg := bytes.Buffer{}
	msg.Write([]byte("hello"))
	// new goroutine to send the message
	go t1.SendMessageToPeer(SendRPC{To: addr2, Payload: &msg})

	received := <-t2.Consume()
	assert.Equal(t, received.From, addr1)
	assert.Equal(t, received.Payload, &msg)
}
