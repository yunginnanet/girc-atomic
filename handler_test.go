package girc

import (
	"log"
	"sync"
	"testing"
)

func TestCaller_AddHandler(t *testing.T) {
	var passChan = make(chan struct{})
	nullClient := &Client{mu: sync.RWMutex{}}
	c := newCaller(nullClient, log.Default())
	c.AddBg("PRIVMSG", func(c *Client, e Event) {
		passChan <- struct{}{}
	})

	go func() {
		c.exec("PRIVMSG", true, nullClient, &Event{})
	}()

	if c.external.lenFor("JONES") != 0 {
		t.Fatalf("wanted %d handlers, got %d", 0, c.internal.lenFor("JONES"))
	}

	if c.external.lenFor("PRIVMSG") != 1 {
		t.Fatalf("wanted %d handlers, got %d", 1, c.external.lenFor("PRIVMSG"))
	}

	<-passChan
}
