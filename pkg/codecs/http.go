package codecs

import (
	"net/http"
	"robot/pkg/common"
	"robot/pkg/logger"
)

type ResponseStatus struct {
	Reason string `json:"reason"`
}

type ResponseValidationErrors struct {
	Validation map[string]string `json:"validationMessages"`
}

// an internal method to be used by other helper functions of codecs
func responseJson(data interface{}, status int, w http.ResponseWriter) http.ResponseWriter {
	byt, err := common.JSON(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		internalErrorJson, _ := common.JSON(&ResponseStatus{
			Reason: "Internal Error",
		})
		_, err = w.Write(internalErrorJson)
	} else {
		_, err = w.Write(byt)
	}

	if err != nil {
		logger.Get().ErrorWithoutSTT("cannot write to output buffer, 500 internal server error returned")
		w.WriteHeader(500)
	}
	return w
}

func ResponseSingleJson(data interface{}, w http.ResponseWriter) http.ResponseWriter {
	return responseJson(data, http.StatusOK, w)
}
func ResponseCreatedJson(data interface{}, w http.ResponseWriter) http.ResponseWriter {
	return responseJson(data, http.StatusCreated, w)
}

func ResponseInternalServerError(customMessage string, w http.ResponseWriter) http.ResponseWriter {
	if customMessage == "" {
		customMessage = "Internal Server Error"
	}
	return responseJson(&ResponseStatus{
		Reason: customMessage,
	}, http.StatusInternalServerError, w)
}

func ResponseStatusJson(customMessage string, status int, w http.ResponseWriter) http.ResponseWriter {
	if customMessage == "" {
		customMessage = http.StatusText(status)
	}
	return responseJson(&ResponseStatus{
		Reason: customMessage,
	}, status, w)
}
