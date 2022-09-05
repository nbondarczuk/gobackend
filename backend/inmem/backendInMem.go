package inmem

type BackendInMem struct {
	kind string
}

type BackendCredentialsInMem struct{}

func (bc BackendCredentialsInMem) ConnectString() string {
	return ""
}

func NewBackendInMem() (BackendInMem, error) {
	return BackendInMem{kind: "inmem"}, nil
}

func (be BackendInMem) Kind() string {
	return be.kind
}

func (be BackendInMem) Version() (string, error) {
	return "irrelevant", nil
}

func (be BackendInMem) Ping() (err error) { return }

func (be BackendInMem) Close() {}
