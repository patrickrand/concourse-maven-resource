package maven

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type Artifact struct {
	Repository string
	GroupID    string
	ArtifactID string
}

func NewArtifact(repository, groupID, artifactID string) *Artifact {
	return &Artifact{
		Repository: repository,
		GroupID:    groupID,
		ArtifactID: artifactID,
	}
}

func (a *Artifact) GetMetadata() (Metadata, error) {
	var metadata Metadata

	uri := fmt.Sprintf("%s/%s/%s/maven-metadata.xml", a.Repository, strings.Replace(a.GroupID, ".", "/", -1), a.ArtifactID)
	resp, err := http.Get(uri)
	if err != nil {
		return metadata, err
	}

	if code := resp.StatusCode; code != http.StatusOK {
		return metadata, fmt.Errorf("%d GET %s", code, uri)
	}

	if err := xml.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return metadata, err
	}
	resp.Body.Close()

	return metadata, nil
}

type Metadata struct {
	XMLName    xml.Name   `xml:"metadata"`
	GroupID    string     `xml:"groupId"`
	ArtifactID string     `xml:"artifactId"`
	Versioning Versioning `xml:"versioning`
}

type Versioning struct {
	XMLName     xml.Name `xml:"versioning"`
	Latest      string   `xml:"latest"`
	Release     string   `xml:"release"`
	Versions    []string `xml:"versions>version"`
	LastUpdated string   `xml:"lastUpdated"`
}
