package bolt

import "errors"

var (
	// ErrBucketNotFound ErrBucketNotFound
	ErrBucketNotFound = errors.New("bucket not found")
	// ErrValueDoesntExist ErrValueDoesntExist
	ErrValueDoesntExist = errors.New("no value against specified key")
)
