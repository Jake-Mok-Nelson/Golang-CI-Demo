package main

import "testing"

func TestHelloWorld(t *testing.T) {
	var msg = helloWorld()
	if msg != "Hello World!" {
		t.Errorf("helloWorld did not work, expected 'Hello World!' but got %s", msg)
	}
}

func TestGoodbyeWorld(t *testing.T) {
	var msg = goodbyeWorld()
	if msg != "Goodbye Cruel World!" {
		t.Errorf("goodbyeWorld did not work, expected 'Goodbye Cruel World!' but got %s", msg)
	}
}
