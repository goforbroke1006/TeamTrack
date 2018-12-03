package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/goforbroke1006/teamtrack/pkg/endpoint"
)

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeCreateTeamHandler creates the handler logic
func makeCreateTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-team", http1.NewServer(endpoints.CreateTeamEndpoint, decodeCreateTeamRequest, encodeCreateTeamResponse, options...))
}

// decodeCreateTeamResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeJoinTeamHandler creates the handler logic
func makeJoinTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/join-team", http1.NewServer(endpoints.JoinTeamEndpoint, decodeJoinTeamRequest, encodeJoinTeamResponse, options...))
}

// decodeJoinTeamResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeJoinTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.JoinTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeJoinTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeJoinTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSetPositionHandler creates the handler logic
func makeSetPositionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/set-position", http1.NewServer(endpoints.SetPositionEndpoint, decodeSetPositionRequest, encodeSetPositionResponse, options...))
}

// decodeSetPositionResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSetPositionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SetPositionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSetPositionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSetPositionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetMatesPositionsHandler creates the handler logic
func makeGetMatesPositionsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-mates-positions", http1.NewServer(endpoints.GetMatesPositionsEndpoint, decodeGetMatesPositionsRequest, encodeGetMatesPositionsResponse, options...))
}

// decodeGetMatesPositionsResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetMatesPositionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetMatesPositionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetMatesPositionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetMatesPositionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
