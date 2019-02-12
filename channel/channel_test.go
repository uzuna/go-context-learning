package channel_test

import (
	"log"
	"testing"

	"github.com/pkg/errors"
)

func TestChannel(t *testing.T) {
	openingch := make(chan struct{}, 1)
	closedch := make(chan struct{})
	close(closedch)

	// Return always (struct{}{}, true) when channel is closed.
	v, ok := <-closedch
	assertTrue(t, v == struct{}{})
	assertTrue(t, !ok)

	select {
	case v, ok := <-openingch:
		log.Println(v, ok)
		t.Logf("This can not be executed")
		t.FailNow()
	default:
		log.Println("Route default. openchannel is waiting")
		openingch <- struct{}{}
	}

	// Return (value, false) when channel is opening
	v, ok = <-openingch
	assertTrue(t, ok)

	// Close channel
	close(openingch)

	// urn always (struct{}{}, true) when channel is closed.
	v, ok = <-openingch
	assertTrue(t, !ok)
	afterClose := openingch
	v, ok = <-afterClose
	assertTrue(t, !ok)
}
func assertTrue(t *testing.T, ok bool) {
	if !ok {
		t.Logf("%+v", errors.Errorf("Must be ok"))
		t.FailNow()
	}
}
