package cli

import "errors"

var (
	ErrWriteWithoutFile = errors.New("cannot specify --write without --file")
	ErrCreateWithID     = errors.New("cannot set id when creating a new item")
	ErrIDRequired       = errors.New("id must be specified")
	ErrIDsWithAll       = errors.New("cannot specify ids when --all is specified")
)
