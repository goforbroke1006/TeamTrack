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

func (l loggingMiddleware) CreateTeam(ctx context.Context, id string, name string) (res bool, err error) {
	defer func() {
		l.logger.Log("method", "CreateTeam", "id", id, "name", name, "res", res, "err", err)
	}()
	return l.next.CreateTeam(ctx, id, name)
}
func (l loggingMiddleware) JoinTeam(ctx context.Context, teamId string, memberId string, deviceInfo string) (res bool, err error) {
	defer func() {
		l.logger.Log("method", "JoinTeam", "teamId", teamId, "memberId", memberId, "deviceInfo", deviceInfo, "res", res, "err", err)
	}()
	return l.next.JoinTeam(ctx, teamId, memberId, deviceInfo)
}

func (l loggingMiddleware) SetPosition(ctx context.Context, data MemberData) (res bool, err error) {
	defer func() {
		l.logger.Log("method", "SetPosition", "data", data, "res", res, "err", err)
	}()
	return l.next.SetPosition(ctx, data)
}
func (l loggingMiddleware) GetMatesPositions(ctx context.Context, memberId string) (res []MemberData, err error) {
	defer func() {
		l.logger.Log("method", "GetMatesPositions", "memberId", memberId, "res", res, "err", err)
	}()
	return l.next.GetMatesPositions(ctx, memberId)
}
