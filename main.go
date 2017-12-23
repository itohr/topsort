package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Service struct {
	DependsOn []string `json:"depends_on"`
}

type Node struct {
	Services map[string]Service
}

func TopSort(services map[string]Service) []string {
	var res []string
	for name, _ := range services {
		fmt.Printf("Check Name: %s\n", name)
		CheckDependency(name, services, &res)
	}
	return res
}

func CheckDependency(name string, services map[string]Service, res *[]string) {
	for _, v := range *res {
		if v == name {
			return
		}
	}
	if services[name].DependsOn != nil {
		for _, dep := range services[name].DependsOn {
			CheckDependency(dep, services, res)
		}
	}
	fmt.Printf("Add Service: %s\n", name)
	*res = append(*res, name)
}

func main() {
	bytes, err := ioutil.ReadFile("services.json")
	if err != nil {
		log.Fatal(err)
	}
	var node Node
	if err := json.Unmarshal(bytes, &node); err != nil {
		log.Fatal(err)
	}

	res := TopSort(node.Services)

	fmt.Printf("Result: %s\n", res)
}
