package main

import (
	"fmt"
	"time"
)

func mySleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println("1")
	mySleep(time.Duration(3) * time.Second)
	fmt.Println("2")
}
