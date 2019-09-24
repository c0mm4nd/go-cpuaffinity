package cpuaffinity

import (
	"crypto/sha256"
	"testing"
	"time"
)

func worker(cpuID int) {
	SetCPUAffinity(cpuID)
	for {
		sha256.Sum256(nil)
	}
}

func TestSetCPUAffinity(t *testing.T) {
	for {
		time.Sleep(time.Second)

		go worker(1)
		go worker(2)
		go worker(3)
		select {}
	}
}
