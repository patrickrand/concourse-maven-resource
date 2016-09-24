package models

type Source struct {
	Repository string `json:"repository"`
	GroupID    string `json:"group_id"`
	ArtifactID string `json:"artifact_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Version struct {
	Number string `json:"number"`
}

func (v Version) String() string {
	return v.Number
}
