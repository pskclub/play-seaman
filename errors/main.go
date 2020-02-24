package errors

import (
	"bitbucket.org/3dsinteractive/seaman"
	"net/http"
)

var (
	NotFound = seaman.Error{
		Status:  http.StatusNotFound,
		Code:    "NOT_FOUND",
		Message: "not found",
	}

	PersisterError = seaman.Error{
		Status:  http.StatusInternalServerError,
		Code:    "PERSISTER_ERROR",
		Message: "persister error",
	}

	IndexerError = seaman.Error{
		Status:  http.StatusInternalServerError,
		Code:    "INDEXER_ERROR",
		Message: "indexer error",
	}

	MQError = seaman.Error{
		Status:  http.StatusInternalServerError,
		Code:    "MQ_ERROR",
		Message: "mq error",
	}

	FatalServerError = seaman.Error{
		Status:  599,
		Code:    "FATAL_SERVER_ERROR",
		Message: "fatal server error",
	}

	InternalServerError = seaman.Error{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "internal server error",
	}

	CacherError = seaman.Error{
		Status:  http.StatusInternalServerError,
		Code:    "CACHER_ERROR",
		Message: "cacher error",
	}

	BadRequest = func(field string, err error) seaman.IError {
		return seaman.Error{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "This " + field + err.Error(),
		}
	}

	JSONFormatIsInvalid = seaman.Error{
		Status:  http.StatusBadRequest,
		Code:    "JSON_FORMAT_IS_INVALID",
		Message: "JSON format is invalid",
	}

	CSVFormatIsInvalid = seaman.Error{
		Status:  http.StatusBadRequest,
		Code:    "CSV_FORMAT_IS_INVALID",
		Message: "CSV format is invalid",
	}

	RequesterError = seaman.Error{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Requester error",
	}
)
