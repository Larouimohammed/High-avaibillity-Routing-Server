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
	wg.Add(2)
	go P.Run(wg)
	go S.Watcher(P, wg)
	for ev := range db.Messages {
		if ev == "change router" {
			wg.Add(2)
			go S.Run(wg)
			go P.Watcher(S, wg)

		}
		defer close(db.Messages)
	}
	wg.Wait()
}
