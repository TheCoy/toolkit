package golib

import "fmt"

const MAX_TASKS = 50

var sem = make(chan int, MAX_TASKS)

type Request struct{
	a,b int
	reply chan int
}

func process(r *Request) {
	fmt.Println("Request resolved:", r.a, r.b)
}

func handle(r *Request) {
	for{
		sem <- 1
		process(r)
		<-sem
	}
}

func service(service chan *Request){
	for{
		request := <-service
		go handle(request)
	}
}

//TestService is an entrance for service.go
func TestService(){
	serv := make(chan *Request)
	go service(serv)
}