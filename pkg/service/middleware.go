package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TeamtrackService) TeamtrackService

type loggingMiddleware struct {
	logger log.Logger
	next   TeamtrackService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TeamtrackService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TeamtrackService) TeamtrackService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Foo(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Foo", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Foo(ctx, s)
}
func (l loggingMiddleware) Bar(ctx context.Context, s string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Bar", "s", s, "e0", e0)
	}()
	return l.next.Bar(ctx, s)
}
func (l loggingMiddleware) Wildfowl(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Wildfowl", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Wildfowl(ctx, s)
}
