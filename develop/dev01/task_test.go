package main

import (
	"testing"
	"time"
)

func Test_getNTPTime(t *testing.T) {
	ntpTime := getNTPTime()
	currTime := time.Now()

	if int(currTime.Sub(ntpTime).Minutes()) > currTime.Second() {
		t.Errorf("Expected - %q, Actual - %q", currTime, ntpTime)
	}
}
