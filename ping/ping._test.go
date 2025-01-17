package ping

import "testing"

func TestPing(t *testing.T) {
	if Pong() != "Pong"{
		t.Error("Expected Pong, got", Pong())
	}
}