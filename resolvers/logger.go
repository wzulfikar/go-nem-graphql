package nemgraphql

import (
	"log"
	"time"
)

func Logger(activity string) func() {
	start := time.Now()
	log.Println("[START] " + activity)
	return func() {
		log.Println("[DONE] "+activity+". Elapsed:", time.Since(start))
	}
}
