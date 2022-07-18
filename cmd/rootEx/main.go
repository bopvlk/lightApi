package main

import (
	"log"

	"7.RootExercise/internal/app/myRootEx"
)

func main() {
	serv := myRootEx.NewRootEx()
	log.Fatal(serv.Start())
	log.Println("Router initalising successfully. Ready to GO!")
}
