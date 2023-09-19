package sample

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSample(t *testing.T) {
	expected := 2
	actual := 3
	if actual != expected {
		t.Errorf("expected: %d, actual %d", expected, actual)
	}
}

func TestVerbose(t *testing.T) {
	t.Log("Only print in verbose mode")
	log.Println("log.println: always printed")
	if testing.Verbose() {
		log.Println("Some verbose, but really useful info")
	}
}

func TestShort(t *testing.T) {
	fmt.Println("Can't be quicker")
}

func TestShortNotSoShort(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skip in short mode")
	}
	time.Sleep(3 * time.Second)
}
