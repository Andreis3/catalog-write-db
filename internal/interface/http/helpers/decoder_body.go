package helpers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andreis3/catalog-write-api/internal/util"
)

func DecoderBodyRequest[T any](req *http.Request) (T, *util.ErrorSchema) {
	defer req.Body.Close()
	var result T
	var jsonUnmarshalTypeError *json.UnmarshalTypeError
	var jsonSyntaxError *json.SyntaxError
	err := json.NewDecoder(req.Body).Decode(&result)
	switch {
	case errors.As(err, &jsonSyntaxError):
		return result, &util.ErrorSchema{
			Code:        "DJ-400",
			Origin:      "DecoderBodyRequest",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json syntax"},
			LogError:    []string{jsonSyntaxError.Error()},
		}
	case errors.As(err, &jsonUnmarshalTypeError):
		return result, &util.ErrorSchema{
			Code:        "DJ-401",
			Origin:      "DecoderBodyRequest",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json field type"},
			LogError:    []string{jsonUnmarshalTypeError.Error()},
		}
	case err != nil:
		return result, &util.ErrorSchema{
			Code:        "DJ-402",
			Origin:      "DecoderBodyRequest",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json"},
			LogError:    []string{err.Error()},
		}
	}
	return result, nil
}
