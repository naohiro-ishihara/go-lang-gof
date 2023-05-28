package main

import (
	"cmd/factory/animals"
	"fmt"
	"log"
)

func main() {
	defer func(){
        if e := recover(); e != nil {
            log.Fatal("Caught error: ", e)
        }
    }()
	
	say := animal.Create(2).Say()

	fmt.Println(say)
}
