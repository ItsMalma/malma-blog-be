package exception

import "fmt"

type RepositoryErrorType string

const (
	RepoErrNotFound RepositoryErrorType = "NOT_FOUND"
)

type RepositoryError struct {
	Type RepositoryErrorType
	Err  string
}

func NewRepoErr(t RepositoryErrorType, err string) RepositoryError {
	return RepositoryError{Type: t, Err: err}
}

func (err RepositoryError) Error() string {
	return fmt.Sprintf("%v: %v", err.Type, err.Err)
}
