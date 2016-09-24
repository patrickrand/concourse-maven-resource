package main

import "github.com/patrickrand/concourse-maven-resource/models"

func main() {

}

type Request struct {
	models.Source  `json:"source"`
	models.Version `json:"version"`
	Params         `json:"params"`
}

type Response struct {
	models.Version `json:"version"`
}

type Params map[string]interface{}
