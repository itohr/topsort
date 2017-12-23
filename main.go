package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Service struct {
	Name      string   `json:"name"`
	DependsOn []string `json:"depends_on"`
}

//func TopSort(services []Service){
//	for _, service := range services {
//	}
//}

func main() {
	bytes, err := ioutil.ReadFile("containers.json")
	if err != nil {
		log.Fatal(err)
	}
	var services []Service
	if err := json.Unmarshal(bytes, &services); err != nil {
		log.Fatal(err)
	}
	for _, p := range services {
		fmt.Printf("%s : %s\n", p.Name, p.DependsOn)
	}
}
