package cpuaffinity_test

import (
	"crypto/sha256"
	"github.com/maoxs2/go-cpuaffinity"
	"time"
)

func Example() {
	for {
		time.Sleep(time.Second)
		go func() {
			cpuaffinity.SetCPUAffinity(10)
			for {
				sha256.Sum256(nil)
			}
		}()
	}
}
