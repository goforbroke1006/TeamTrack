package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/goforbroke1006/teamtrack/pkg/service"
)

type CreateTeamRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateTeamResponse struct {
	Res bool  `json:"res"`
	Err error `json:"err"`
}

func MakeCreateTeamEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTeamRequest)
		res, err := s.CreateTeam(ctx, req.Id, req.Name)
		return CreateTeamResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r CreateTeamResponse) Failed() error {
	return r.Err
}

type JoinTeamRequest struct {
	TeamId     string `json:"team_id"`
	MemberId   string `json:"member_id"`
	DeviceInfo string `json:"device_info"`
}

type JoinTeamResponse struct {
	Res bool  `json:"res"`
	Err error `json:"err"`
}

func MakeJoinTeamEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(JoinTeamRequest)
		res, err := s.JoinTeam(ctx, req.TeamId, req.MemberId, req.DeviceInfo)
		return JoinTeamResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r JoinTeamResponse) Failed() error {
	return r.Err
}

type SetPositionRequest struct {
	Data service.MemberData `json:"data"`
}

type SetPositionResponse struct {
	Res bool  `json:"res"`
	Err error `json:"err"`
}

func MakeSetPositionEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetPositionRequest)
		res, err := s.SetPosition(ctx, req.Data)
		return SetPositionResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r SetPositionResponse) Failed() error {
	return r.Err
}

type GetMatesPositionsRequest struct {
	MemberId string `json:"member_id"`
}

type GetMatesPositionsResponse struct {
	Res []service.MemberData `json:"res"`
	Err error                `json:"err"`
}

func MakeGetMatesPositionsEndpoint(s service.TeamtrackService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMatesPositionsRequest)
		res, err := s.GetMatesPositions(ctx, req.MemberId)
		return GetMatesPositionsResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r GetMatesPositionsResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) CreateTeam(ctx context.Context, id string, name string) (res bool, err error) {
	request := CreateTeamRequest{
		Id:   id,
		Name: name,
	}
	response, err := e.CreateTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateTeamResponse).Res, response.(CreateTeamResponse).Err
}

func (e Endpoints) JoinTeam(ctx context.Context, teamId string, memberId string, deviceInfo string) (res bool, err error) {
	request := JoinTeamRequest{
		DeviceInfo: deviceInfo,
		MemberId:   memberId,
		TeamId:     teamId,
	}
	response, err := e.JoinTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(JoinTeamResponse).Res, response.(JoinTeamResponse).Err
}

func (e Endpoints) SetPosition(ctx context.Context, data service.MemberData) (res bool, err error) {
	request := SetPositionRequest{Data: data}
	response, err := e.SetPositionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SetPositionResponse).Res, response.(SetPositionResponse).Err
}

func (e Endpoints) GetMatesPositions(ctx context.Context, memberId string) (res []service.MemberData, err error) {
	request := GetMatesPositionsRequest{MemberId: memberId}
	response, err := e.GetMatesPositionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetMatesPositionsResponse).Res, response.(GetMatesPositionsResponse).Err
}
