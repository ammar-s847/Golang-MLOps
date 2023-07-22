package types

type Model struct {
	id      float64
	name    string
	version string
	urlType string
	url     string
	owner   User
}

func (model *Model) Version() string {
	return model.version
}

func (model *Model) UpdateVersion(version string) {
	model.version = version
	// Change in DB and Cache
}
