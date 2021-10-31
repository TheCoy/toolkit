package toy

import "fmt"

func G() {
	defer func() {
		fmt.Println("defer: func ended")
	}()
	r := F()
	fmt.Printf("func G executed:%v", r)
}

func F() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer: panic catched:", err)
		}
	}()
	panic("a")
	return 3.14
}
