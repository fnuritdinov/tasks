package errors

import "errors"

var ErrFromValidate = errors.New("error from validate")
var ErrNotFound = errors.New("not found")
var ErrBadRequest = errors.New("bad request")
var ErrInternal = errors.New("internal error")
