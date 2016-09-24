package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/patrickrand/concourse-maven-resource/maven"
	"github.com/patrickrand/concourse-maven-resource/models"
)

func main() {
	resp, err := Check(os.Stdin)
	if err != nil {
		exit(err.Error())
	}

	if err := json.NewEncoder(os.Stdout).Encode(resp); err != nil {
		exit(err.Error())
	}
}

func Check(r io.Reader) (Response, error) {
	var resp Response

	var req Request
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return resp, errors.New("failed to decode incoming check request: " + err.Error())
	}

	artifact := maven.NewArtifact(
		req.Source.Repository,
		req.Source.GroupID,
		req.Source.ArtifactID,
	)

	metadata, err := artifact.GetMetadata()
	if err != nil {
		return resp, errors.New("failed to fetch artifact metadata: " + err.Error())
	}

	for _, v := range metadata.Versioning.Versions {
		resp = append(resp, models.Version{Number: v})
	}

	return resp, nil
}

func exit(msg string) {
	fmt.Println("error: " + msg)
	os.Exit(1)
}

type Request struct {
	models.Source  `json:"source"`
	models.Version `json:"version"`
}

type Response []models.Version
