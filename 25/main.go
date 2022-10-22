// Реализовать собственную функцию sleep.

package main

import "time"

func Sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	Sleep(5 * time.Second)
}
