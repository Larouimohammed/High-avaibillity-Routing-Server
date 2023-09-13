package main

import (
	db "HA/DB"
	"HA/routerplugin"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	P := routerplugin.NewRouter("R1", "primary", true, []byte{192, 168, 1, 12}, []byte{255, 255, 255, 0}, "enp")
	S := routerplugin.NewRouter("R2", "Secondary", false, []byte{192, 168, 1, 13}, []byte{255, 255, 255, 0}, "enp")
	P.Run(wg)
	S.Watcher(P, wg)
	for ev := range db.Messages {

		if ev == P.Key {
			S.Run(wg)
			P.Watcher(S, wg)

		}
		if ev == S.Key {
			P.Run(wg)
			S.Watcher(P, wg)

		}
		defer wg.Done()
		defer close(db.Messages)
	}
	wg.Wait()
}
