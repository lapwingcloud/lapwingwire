package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

type Responder interface {
	OK() *Response
	BadRequest() *Response
	Unauthorized() *Response
	Forbidden() *Response
	NotFound() *Response
	Conflict() *Response
	InternalServerError() *Response
	NotImplemented() *Response
}

type responder struct {
	logger *zerolog.Logger
	w      http.ResponseWriter
}

func NewResponder(w http.ResponseWriter, r *http.Request) Responder {
	return &responder{logger: hlog.FromRequest(r), w: w}
}

func (t *responder) OK() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusOK,
	}
}

func (t *responder) BadRequest() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusBadRequest,
	}
}

func (t *responder) Unauthorized() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusUnauthorized,
	}
}

func (t *responder) Forbidden() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusForbidden,
	}
}

func (t *responder) NotFound() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusNotFound,
	}
}

func (t *responder) Conflict() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusConflict,
	}
}

func (t *responder) InternalServerError() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusInternalServerError,
	}
}

func (t *responder) NotImplemented() *Response {
	return &Response{
		logger: t.logger,
		w:      t.w,
		status: http.StatusNotImplemented,
	}
}

type Response struct {
	logger  *zerolog.Logger
	w       http.ResponseWriter
	status  int
	err     error
	message string
	data    any
}

func (t *Response) Err(err error) *Response {
	t.err = err
	return t
}

func (t *Response) Msg(message string) {
	t.message = message
	t.write()
}

func (t *Response) Data(data any) {
	t.data = data
	t.write()
}

func (t *Response) write() {
	if t.status >= 500 {
		t.logger.Error().Err(t.err).Msg(t.message)
	} else if t.status >= 400 {
		t.logger.Info().Err(t.err).Msg(t.message)
	}
	r := struct {
		Status  int    `json:"status,omitempty"`
		Error   string `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
		Data    any    `json:"data,omitempty"`
	}{
		Status:  t.status,
		Message: t.message,
		Data:    t.data,
	}
	if t.err != nil {
		r.Error = t.err.Error()
	}
	bytes, err := json.Marshal(r)
	t.w.Header().Set("Content-Type", "application/json")
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to encode response object to json")
		t.w.WriteHeader(500)
		if _, err = t.w.Write([]byte(fmt.Sprintf(`{"status":500,"error":"%v","message":"failed to encode response object to json"}`, err))); err != nil {
			t.logger.Error().Err(err).Msg("failed to write response")
		}
		return
	}
	t.w.WriteHeader(t.status)
	if _, err = t.w.Write(bytes); err != nil {
		t.logger.Error().Err(err).Msg("failed to write response")
	}
}
