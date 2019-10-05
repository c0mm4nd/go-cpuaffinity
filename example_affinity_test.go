package cpuaffinity_test

import (
	"crypto/sha256"
	"github.com/maoxs2/go-cpuaffinity"
)

func Example() {
	go func() {
		cpuaffinity.SetCPUAffinity(0)
		for {
			sha256.Sum256(nil)
		}
	}()

	go func() {
		cpuaffinity.SetCPUAffinity(2)
		for {
			sha256.Sum256(nil)
		}
	}()

	go func() {
		cpuaffinity.SetCPUAffinity(3)
		for {
			sha256.Sum256(nil)
		}
	}()

	select {}
}
