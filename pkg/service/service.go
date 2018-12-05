package service

import (
	"context"
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/goforbroke1006/teamtrack/pkg/domain"
	"github.com/goforbroke1006/teamtrack/pkg/entity"
)

// TeamtrackService describes the service.
type TeamtrackService interface {
	CreateTeam(ctx context.Context, id string, name string, from, till time.Time) (res bool, err error)
	JoinTeam(ctx context.Context, teamId, memberId, deviceInfo string) (res bool, err error)
	SetPosition(ctx context.Context, data domain.MemberData) (res bool, err error)
	GetMatesPositions(ctx context.Context, memberId string) (res []domain.MemberData, err error)
}
type basicTeamtrackService struct {
	db *gorm.DB
}

// NewBasicTeamtrackService returns a naive, stateless implementation of TeamtrackService.
func NewBasicTeamtrackService(db *gorm.DB) TeamtrackService {
	return &basicTeamtrackService{
		db: db,
	}
}

// New returns a TeamtrackService with all of the expected middleware wired in.
func New(middleware []Middleware, db *gorm.DB) TeamtrackService {
	var svc TeamtrackService = NewBasicTeamtrackService(db)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicTeamtrackService) CreateTeam(ctx context.Context,
	id string, name string, from, till time.Time,
) (res bool, err error) {
	if 0 == len(id) || 0 == len(name) {
		return false, errors.New("team data required")
	}
	if till.Unix() < time.Now().Unix() {
		return false, errors.New("non-actual event")
	}
	if till.Unix() <= from.Unix() {
		return false, errors.New("wrong time range")
	}

	team := entity.Team{ID: id, Name: name, ActiveFrom: from, ActiveTill: till}
	b.db.Create(&team)
	res = !b.db.NewRecord(team)
	return res, err
}
func (b *basicTeamtrackService) JoinTeam(ctx context.Context, teamId string, memberId string, deviceInfo string) (res bool, err error) {
	// TODO implement the business logic of JoinTeam
	return res, err
}
func (b *basicTeamtrackService) SetPosition(ctx context.Context, info domain.MemberData) (res bool, err error) {
	// TODO implement the business logic of SetPosition
	return res, err
}
func (b *basicTeamtrackService) GetMatesPositions(ctx context.Context, memberId string) (res []domain.MemberData, err error) {
	// TODO implement the business logic of GetMatesPositions
	return res, err
}
