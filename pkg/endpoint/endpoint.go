package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/goforbroke1006/teamtrack/pkg/service"
)

// FooRequest collects the request parameters for the Foo method.
type FooRequest struct {
	S string `json:"s"`
}

// FooResponse collects the response parameters for the Foo method.
type FooResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeFooEndpoint returns an endpoint that invokes Foo on the service.
func MakeFooEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		rs, err := s.Foo(ctx, req.S)
		return FooResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r FooResponse) Failed() error {
	return r.Err
}

// BarRequest collects the request parameters for the Bar method.
type BarRequest struct {
	S string `json:"s"`
}

// BarResponse collects the response parameters for the Bar method.
type BarResponse struct {
	E0 error `json:"e0"`
}

// MakeBarEndpoint returns an endpoint that invokes Bar on the service.
func MakeBarEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BarRequest)
		e0 := s.Bar(ctx, req.S)
		return BarResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r BarResponse) Failed() error {
	return r.E0
}

// WildfowlRequest collects the request parameters for the Wildfowl method.
type WildfowlRequest struct {
	S string `json:"s"`
}

// WildfowlResponse collects the response parameters for the Wildfowl method.
type WildfowlResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeWildfowlEndpoint returns an endpoint that invokes Wildfowl on the service.
func MakeWildfowlEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WildfowlRequest)
		rs, err := s.Wildfowl(ctx, req.S)
		return WildfowlResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r WildfowlResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Foo implements Service. Primarily useful in a client.
func (e Endpoints) Foo(ctx context.Context, s string) (rs string, err error) {
	request := FooRequest{S: s}
	response, err := e.FooEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FooResponse).Rs, response.(FooResponse).Err
}

// Bar implements Service. Primarily useful in a client.
func (e Endpoints) Bar(ctx context.Context, s string) (e0 error) {
	request := BarRequest{S: s}
	response, err := e.BarEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BarResponse).E0
}

// Wildfowl implements Service. Primarily useful in a client.
func (e Endpoints) Wildfowl(ctx context.Context, s string) (rs string, err error) {
	request := WildfowlRequest{S: s}
	response, err := e.WildfowlEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WildfowlResponse).Rs, response.(WildfowlResponse).Err
}
