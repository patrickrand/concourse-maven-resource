package maven

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func MetadataMockServer(t *testing.T, groupID, artifactID, version string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadata := Metadata{
			GroupID:    groupID,
			ArtifactID: artifactID,
			Versioning: Versioning{
				Latest:   version,
				Release:  version,
				Versions: []string{version},
			},
		}
		if err := xml.NewEncoder(w).Encode(metadata); err != nil {
			t.Fatal("Failed to marshal xml response", err)
		}
	}))
}

func TestDownloadLatest(t *testing.T) {
	tests := []struct {
		repo, group, artifact, version string
		error
	}{
		{repo: "", group: "com.foo", artifact: "bar", version: "1.0"},
	}

	for _, tc := range tests {
		server := MetadataMockServer(t, tc.group, tc.artifact, tc.version)
		defer server.Close()
		artifact := NewArtifact(server.URL, tc.group, tc.artifact)
		if err := artifact.DownloadLatest(os.TempDir() + "/foo"); err != tc.error {
			t.Errorf("expected: %v, got: %v", tc.error, err)
		}
	}
}

func TestGetMetadata(t *testing.T) {
	tests := []struct {
		repo, group, artifact, version string
		error
	}{
		{repo: "", group: "com.foo", artifact: "bar", version: "1.0"},
	}

	for _, tc := range tests {
		server := MetadataMockServer(t, tc.group, tc.artifact, tc.version)
		defer server.Close()
		artifact := NewArtifact(server.URL, tc.group, tc.artifact)
		metadata, err := artifact.GetMetadata()
		if err != tc.error {
			t.Errorf("expected: %v, got: %v", tc.error, err)
		}

		t.Logf("%#v, %#v", metadata, err)
	}
}
