package maven

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

func (a *Artifact) DownloadLatest(dest string) error {
	metadata, err := a.GetMetadata()
	if err != nil {
		return err
	}

	version := metadata.Versioning.Latest

	uri := fmt.Sprintf("%s/%s/%s/%s/%s-%s.jar", a.Repository, strings.Replace(a.GroupID, ".", "/", -1), a.ArtifactID, version, a.ArtifactID, version)
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if err := ioutil.WriteFile(dest, data, 0777); err != nil {
		return err
	}

	return nil
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
