package maven

import "testing"

func TestDownloadLatest(t *testing.T) {
	tests := []struct {
		repo, group, artifact string
		error
	}{
		{repo: "", group: "", artifact: ""},
	}

	for _, tc := range tests {
		artifact := NewArtifact(tc.repo, tc.group, tc.artifact)
		if err := artifact.DownloadLatest(""); err != tc.error {
			t.Errorf("expected: %v, got: %v", tc.error, err)
		}
	}
}

func TestGetMetadata(t *testing.T) {
	tests := []struct {
		repo, group, artifact string
		error
	}{
		{repo: "", group: "", artifact: ""},
	}

	for _, tc := range tests {
		artifact := NewArtifact(tc.repo, tc.group, tc.artifact)
		metadata, err := artifact.GetMetadata()
		if err != tc.error {
			t.Errorf("expected: %v, got: %v", tc.error, err)
		}

		t.Logf("%#v, %#v", metadata, err)
	}
}
