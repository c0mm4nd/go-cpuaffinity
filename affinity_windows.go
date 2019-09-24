// Package cpuaffinity
// affinity will combine the task(goroutine) with cpu, no longer schedule , which speed up the task
package cpuaffinity

import (
	"runtime"
)

/*
#include <Windows.h>

void lock(int cpuID) {
    HANDLE process = GetCurrentThread();
    DWORD_PTR processAffinityMask = 1 << cpuID;

    SetThreadAffinityMask(process, processAffinityMask);
}
*/
import "C"

// SetCPUAffinity locks the goroutine to the thread and then locks the thread to the CPU
func SetCPUAffinity(cpuID int) {
	runtime.LockOSThread()
	C.lock(C.int(cpuID))
}
