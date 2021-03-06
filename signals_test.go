package abutil

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestOnSignal(t *testing.T) {
	done := make(chan bool)
	sg := syscall.SIGHUP

	go OnSignal(func(s os.Signal) {
		if s != sg {
			t.Errorf("Expected signal %s, but got %s", sg, s)
		}

		done <- true
	})

	// Send interrupt after 10ms
	time.AfterFunc(10*time.Millisecond, func() {
		syscall.Kill(syscall.Getpid(), sg)
	})
	<-done
}

func ExampleOnSignal() {
	done := make(chan bool)

	go OnSignal(func(s os.Signal) {
		fmt.Printf("Got signal %s\n", s)

		done <- true
	})

	// Emulate SIGINT
	time.AfterFunc(10*time.Millisecond, func() {
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})

	// Wait until we got our signal
	<-done

	// Output: Got signal interrupt
}
