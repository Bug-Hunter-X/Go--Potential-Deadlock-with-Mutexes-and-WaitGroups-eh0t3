# Go: Potential Deadlock with Mutexes and WaitGroups

This repository demonstrates a subtle deadlock that can occur in Go programs using goroutines, mutexes, and wait groups.  The program, as written, *may* deadlock; it's race-dependent.

## The Bug

The `bug.go` file contains a program that creates multiple goroutines, each acquiring a mutex (`m.Lock()`) before printing a message and releasing the mutex (`m.Unlock()`).  The `sync.WaitGroup` is used to ensure all goroutines complete before the main function exits.  If the simulated work is removed, the program's chances of deadlocking increase because the scheduler's behavior influences the timing of mutex acquisition and release.