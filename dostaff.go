package main

import (
	"time"
)

func repeatme(fn func(s int)) chan bool {
	ticker := time.NewTicker(2 * time.Second)
	t := 5
	quit := make(chan bool)
	// log.Print("waiting for 'for loop to start'")
	go func() {
		for {
			select {
			case <-ticker.C:
				// log.Print("Executing function in repeateme()")
				fn(t)
				t++
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return quit
}
