package myRootEx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"7.RootExercise/models"
)

type Message struct {
	Message string `json:"message"`
}

var rootIn models.RootIn

func (roo *RootEX) RootsOfQuadraticEquation() (*models.RootOut, error) {
	var err error
	/////
	JsonRead()
	/////
	fmt.Println(rootIn)
	d := rootIn.B*rootIn.B - 4*rootIn.A*rootIn.C
	if (rootIn.A == 0 || rootIn.B == 0 && rootIn.C == 0) || (d == 0) {
		roo.rOut.Nroots = 1
	} else if (rootIn.B == 0 && rootIn.A/rootIn.C < 0) || (d < 0) {
		roo.rOut.Nroots = 0
	} else if (rootIn.B == 0 && rootIn.A/rootIn.C > 0) || (rootIn.C == 0) || (d < 0) {
		roo.rOut.Nroots = 2
	} else {
		return nil, err
	}
	roo.rOut.A, roo.rOut.B, roo.rOut.C = rootIn.A, rootIn.B, rootIn.C

	return roo.rOut, nil
}

func (roo *RootEX) JsonRead() {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		roo.l.Fatalln("Неможливо відкрити json файл")
	}
	defer jsonFile.Close()

	roo.l.Println("File descriptor successefuly created")

	byteVaule, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		roo.l.Fatalln("Неможливо зчитати json файл")
	}

	json.Unmarshal(byteVaule, &rootIn)
	if err != nil {
		log.Fatal(err)
	}
}
