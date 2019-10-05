// Package cpuaffinity
// affinity will combine the task(goroutine) with cpu, no longer schedule , which speed up the task
package cpuaffinity

import (
	"runtime"
)

/*
#include <Windows.h>

void lock(int cpuID) {
    HANDLE thread = GetCurrentThread();
    DWORD_PTR threadAffinityMask = 1 << cpuID;

    SetThreadAffinityMask(thread, threadAffinityMask);
}

void lockMask(unsigned long long mask) {
    HANDLE thread = GetCurrentThread();
    DWORD_PTR threadAffinityMask = mask;

    SetThreadAffinityMask(thread, threadAffinityMask);
}
*/
import "C"

// SetCPUAffinity locks the goroutine to the thread and then locks the thread to the CPU
func SetCPUAffinity(cpuID int) {
	runtime.LockOSThread()
	C.lock(C.int(cpuID))
}

func SetCPUAffinityMask(mask uint64) {
	runtime.LockOSThread()
	C.lockMask(C.ulonglong(mask))
}
