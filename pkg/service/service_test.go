package service

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/goforbroke1006/teamtrack/pkg/domain"
	"github.com/goforbroke1006/teamtrack/pkg/entity"
)

func createTestMiddleWares() []Middleware {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var mw []Middleware
	mw = append(mw, LoggingMiddleware(logger))
	return mw
}

func createTestDB() *gorm.DB {
	db, _ := gorm.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	db.AutoMigrate(
		entity.Team{},
		entity.Member{},
		entity.Location{},
	)
	return db
}

func TestCreateTeamAction(t *testing.T) {
	svc := New(createTestMiddleWares(), createTestDB())

	var FakeID = fmt.Sprintf("test-%d", time.Now().Unix())
	_, err := svc.CreateTeam(context.Background(), FakeID, "qwer", time.Unix(0, 0), time.Now())
	if nil != err {
		t.Error(err)
	}

	_, err = svc.CreateTeam(context.Background(), FakeID, "asdf", time.Unix(0, 0), time.Now())
	if nil == err {
		t.Error(errors.New("should init error about existent team"))
	}

	_, err = svc.CreateTeam(context.Background(), FakeID, "asdf", time.Now(), time.Now())
	if nil == err {
		t.Error(errors.New("should return error about wrong time range"))
	}
}

const MemberOne = "Blade"
const MemberTwo = "Abraham Whistler"

func createBladeTeam(t *testing.T, svc TeamtrackService) {
	var FakeID = "vampire-hunter-vacation-2018"
	_, err := svc.CreateTeam(context.Background(), FakeID, "Cool Team # 1", time.Unix(0, 0), time.Now())
	if nil != err {
		t.Error(err)
	}

	_, err = svc.JoinTeam(context.Background(), FakeID, MemberOne, "Android 4.5.6")
	if nil != err {
		t.Error(err)
	}

	_, err = svc.JoinTeam(context.Background(), FakeID, MemberTwo, "Android 4.5.6")
	if nil != err {
		t.Error(err)
	}

	_, err = svc.JoinTeam(context.Background(), FakeID, "Hannibal King", "Android 1.5.6")
	if nil != err {
		t.Error(err)
	}
}

func TestJoinTeam(t *testing.T) {
	svc := New(createTestMiddleWares(), createTestDB())

	createBladeTeam(t, svc)

	res, err := svc.GetMatesPositions(context.Background(), MemberOne)
	if nil != err {
		t.Error(err)
	}

	assert.Equal(t, 2, len(res), "wrong teammates count")
}

func TestGetMatesPositions(t *testing.T) {
	svc := New(createTestMiddleWares(), createTestDB())

	createBladeTeam(t, svc)

	b, e := svc.SetPosition(context.Background(), domain.MemberData{
		MemberId: MemberTwo,
		Lat:      10,
		Lng:      20,
	})
	if nil != e {
		t.Error(e)
	}
	assert.Equal(t, true, b, "should works")

	res, err := svc.GetMatesPositions(context.Background(), MemberOne)
	if nil != err {
		t.Error(err)
	}

	assert.Equal(t, 2, len(res), "wrong teammates count")

	var pos *domain.MemberData
	for _, m := range res {
		if m.MemberId == MemberTwo {
			pos = &m
		}
	}

	assert.NotEqual(t, nil, pos, "member one should see member two")

	if nil != pos {
		assert.Equal(t, 20, pos.Lng, "member two must be in correct location")
	}

}
