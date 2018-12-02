package service

import "context"

// TeamtrackService describes the service.
type TeamtrackService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
	Bar(ctx context.Context, s string) error
	Wildfowl(ctx context.Context, s string) (rs string, err error)
}
type basicTeamtrackService struct{}

func (b *basicTeamtrackService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}
func (b *basicTeamtrackService) Bar(ctx context.Context, s string) (e0 error) {
	// TODO implement the business logic of Bar
	return e0
}
func (b *basicTeamtrackService) Wildfowl(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Wildfowl
	rs = "Hello!"
	return rs, err
}

// NewBasicTeamtrackService returns a naive, stateless implementation of TeamtrackService.
func NewBasicTeamtrackService() TeamtrackService {
	return &basicTeamtrackService{}
}

// New returns a TeamtrackService with all of the expected middleware wired in.
func New(middleware []Middleware) TeamtrackService {
	var svc TeamtrackService = NewBasicTeamtrackService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
