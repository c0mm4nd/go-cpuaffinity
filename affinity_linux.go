package cpuaffinity

import (
	"runtime"
)

/*
#define _GNU_SOURCE
#include <sched.h>
#include <pthread.h>

void lock(int cpuid) {
    pthread_t tid;
    cpu_set_t cpuset;

    tid = pthread_self();
    CPU_ZERO(&cpuset);
    CPU_SET(cpuid, &cpuset);
    pthread_setaffinity_np(tid, sizeof(cpu_set_t), &cpuset);
}

void lockMask(unsigned long long mask) {
    pthread_t tid;
    cpu_set_t cpuset;

    tid = pthread_self();
    CPU_ZERO(&cpuset);
    for (int i = 0; ; ++i)
    {
        if ( mask >> i == 0 )
        {
            break;
        }
        if ( (mask >> i) & (1<<i) )
        {
            CPU_SET(i, &cpuset);
        }
    }

    pthread_setaffinity_np(tid, sizeof(cpu_set_t), &cpuset);
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
