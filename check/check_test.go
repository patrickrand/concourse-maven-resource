package main

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
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
				models.Source{"", "group", "artifact", "username", "password"},
				models.Version{"0.0.0"},
			},
			Response: Response{
				{"0.0.0"},
			},
		},
	}

	for _, tc := range tests {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			metadata := maven.Metadata{
				GroupID:    tc.GroupID,
				ArtifactID: tc.ArtifactID,
				Versioning: maven.Versioning{
					Latest:   tc.Version.Number,
					Release:  tc.Version.Number,
					Versions: []string{tc.Version.Number},
				},
			}
			if err := xml.NewEncoder(w).Encode(metadata); err != nil {
				t.Fatal("Failed to marshal xml response", err)
			}
		}))
		defer mockServer.Close()
		artifact := maven.NewArtifact(mockServer.URL, tc.Request.Source.GroupID, tc.Request.Source.ArtifactID)
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
