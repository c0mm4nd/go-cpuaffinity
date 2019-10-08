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

func workerMask(mask uint64) {
	SetCPUAffinityMask(mask)
	for {
		sha256.Sum256(nil)
	}
}

func TestSetCPUAffinity(t *testing.T) {
	time.Sleep(time.Second)

	cpus := []int{0, 1, 3}
	for _, cpuid := range cpus {
		go worker(cpuid)
	}

	select {}
}

func TestSetCPUAffinityMask(t *testing.T) {
	time.Sleep(time.Second)

	go workerMask(1<<0 | 1<<1 | 1<<3)
	go workerMask(1<<0 | 1<<1 | 1<<3)
	go workerMask(1<<0 | 1<<1 | 1<<3)

	select {}
}
