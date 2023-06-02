package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type Response struct {
	w http.ResponseWriter
	r *http.Request

	Status     int    `json:"status,omitempty"`
	Error      string `json:"error,omitempty"`   // application-level error message, for debugging
	Message    string `json:"message,omitempty"` // application-level error message, for debugging
	DataObject any    `json:"data,omitempty"`
}

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{w: w, r: r}
}

func (t *Response) OK() *Response {
	t.Status = http.StatusOK
	return t
}

func (t *Response) BadRequest() *Response {
	t.Status = http.StatusBadRequest
	return t
}

func (t *Response) Unauthorized() *Response {
	t.Status = http.StatusUnauthorized
	return t
}

func (t *Response) Forbidden() *Response {
	t.Status = http.StatusForbidden
	return t
}

func (t *Response) Conflict() *Response {
	t.Status = http.StatusConflict
	return t
}

func (t *Response) InternalServerError() *Response {
	t.Status = http.StatusInternalServerError
	return t
}

func (t *Response) NotImplemented() *Response {
	t.Status = http.StatusNotImplemented
	return t
}

func (t *Response) Err(err error) *Response {
	t.Error = err.Error()
	return t
}

func (t *Response) Msg(message string) {
	t.Message = message
	t.write()
}

func (t *Response) Data(data any) {
	t.DataObject = data
	t.write()
}

func (t *Response) write() {
	logger := hlog.FromRequest(t.r)
	t.w.WriteHeader(t.Status)
	t.w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(*t)
	if err != nil {
		logger.Error().Err(err).Msg("failed to encode response object to json")
		if _, err = t.w.Write([]byte(fmt.Sprintf(`{"status":500,"error":"%v","message":"failed to encode response object to json"}`, err))); err != nil {
			logger.Error().Err(err).Msg("failed to write response")
		}
	}
	if _, err = t.w.Write(bytes); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
	}
}
