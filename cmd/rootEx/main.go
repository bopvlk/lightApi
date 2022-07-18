package main

import (
	"log"

	"7.RootExercise/internal/app/myRootEx"
)

func main() {
	serv := myRootEx.NewRootEx()
	serv.Start()
	log.Println("Router initalising successfully. Ready to GO!")
}
