package myRootEx

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"7.RootExercise/models"
)

func initHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (roo *RootEX) Grab(w http.ResponseWriter, r *http.Request) {
	initHeader(w)
	roo.l.Println("Get infos about all numbers")
	var rootsIn *models.RootIn

	err := json.NewDecoder(r.Body).Decode(&rootsIn)
	if err != nil {
		roo.l.Println("Invalid  json resived from client")
		m := Message{
			Message: "Provided json is invalid",
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(m)
	}

	byteArr, err := json.MarshalIndent(rootsIn, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("output.json", byteArr, 0666)
	if err != nil {
		roo.l.Fatal("не зміг записати в JSON", err)
	}

	roo.l.Println("data was recorded")
	m := Message{
		Message: "All good!",
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(m)
}

func (roo *RootEX) Solve(w http.ResponseWriter, r *http.Request) {
	initHeader(w)
	roo.l.Println("Tru the root of a quadratic equation")
	rOut, err := roo.RootsOfQuadraticEquation()
	fmt.Println("|||||||||||", rOut, "|||||||||||")
	if err != nil {
		roo.l.Println("Troble with RootsOfQuadraticEquation()")
		m := Message{
			Message: "Som troble with data",
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(m)
	}
	roo.l.Println("All good")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(rOut)
}
