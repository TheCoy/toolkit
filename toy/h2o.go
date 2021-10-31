package toy

import (
	"fmt"
	"time"
)

var container map[string]int

func hydrogen() {
	for i := 0; i < 100; i++ {
		lock.Lock()
		container["h"]++
		lock.Unlock()
	}
}

func oxygen() {
	for i := 0; i < 40; i++ {
		lock.Lock()
		container["o"]++
		lock.Unlock()
	}
}

func synthesizeH2O() {
	for {
		lock.Lock()
		if container["h"] >= 2 && container["o"] >= 1 {
			container["h"] -= 2
			container["o"] -= 1
			fmt.Println(">>>> H2O")
		}

		lock.Unlock()
	}
}

func RunSynthesizeH2O() {
	container = make(map[string]int)
	go hydrogen()
	go oxygen()
	go synthesizeH2O()

	time.Sleep(2*time.Second)
}
