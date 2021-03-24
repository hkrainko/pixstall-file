package model

type FileACL struct {
	ID    string
	Owner string
	ACL   []string
	State FileState
}

type FileState string

const (
	FileStateActive  = "A"
	FileStateHidden  = "H"
	FileStateDeleted = "D"
)
