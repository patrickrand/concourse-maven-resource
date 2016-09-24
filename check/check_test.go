package main

import (
	"testing"

	"github.com/patrickrand/concourse-maven-resource/maven"
	"github.com/patrickrand/concourse-maven-resource/models"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		Request
		Response
		error
	}{
		{
			Request: Request{
				models.Source{"repo", "group", "artifact", "username", "password"},
				models.Version{"0.0.0"},
			},
			Response: Response{
				{"0.0.0"},
			},
		},
	}

	for _, tc := range tests {
		artifact := maven.NewArtifact(tc.Request.Source.Repository, tc.Request.Source.GroupID, tc.Request.Source.ArtifactID)
		metadata, err := artifact.GetMetadata()
		if err != tc.error {
			t.Errorf("expected: %v, got: %v", tc.error, err)
		}

		versions := metadata.Versioning.Versions
		for i, resp := range tc.Response {
			if n := len(versions); i >= n {
				t.Errorf("expected: %s, got: <nil>", resp)
				continue
			}
			if expected, got := tc.Response[i].String(), versions[i]; expected != got {
				t.Errorf("expected: %s, got: %s", expected, got)
			}
		}
	}
}
