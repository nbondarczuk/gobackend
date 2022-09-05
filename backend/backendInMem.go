package backend

type BackendInMem struct {
	BackendCommonInfo
}

type BackendCredentialsInMem struct{}

func NewBackendInMem() (BackendInMem, error) {
	return BackendInMem{BackendCommonInfo{"inmem"}}, nil
}

func (be BackendInMem) Kind() string {
	return be.kind
}

func (be BackendInMem) Version() (string, error) {
	return "irrelevant", nil
}

func (be BackendInMem) Ping() (err error) { return }

func (be BackendInMem) Close() {}
