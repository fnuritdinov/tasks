package handlers

import (
	errs "TaskManager/pkg/errors"
	"net/http"

	"go.uber.org/zap"
)

func handleError(w http.ResponseWriter, log *zap.Logger, err error) {
	switch err {
	case errs.ErrBadRequest:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errs.ErrNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	case errs.ErrInternal:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case errs.ErrFromValidate:
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return
}
