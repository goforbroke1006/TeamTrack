package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/goforbroke1006/teamtrack/pkg/entity"
	"github.com/goforbroke1006/teamtrack/pkg/service"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"os"
	"testing"
	"time"
)
import _ "github.com/mattn/go-sqlite3"

func createTestDB() *gorm.DB {
	db, _ := gorm.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	//db, _ := gorm.Open("sqlite3", "file:test.db?mode=memory")
	//db, _ := gorm.Open("sqlite3", "file:test.db?")
	db.AutoMigrate(
		entity.Team{},
		entity.Member{},
		entity.Location{},
	)
	return db
}

func TestCreateTeamAction(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	db := createTestDB()
	svc := service.New(getServiceMiddleware(logger), db)
	var FakeID = fmt.Sprintf("test-%d", time.Now().Unix())
	_, err := svc.CreateTeam(context.Background(), FakeID, "qwer", time.Unix(0, 0), time.Now())
	if nil != err {
		t.Error(err)
	}
	_, err = svc.CreateTeam(context.Background(), FakeID, "asdf", time.Unix(0, 0), time.Now())
	if nil == err {
		t.Error(errors.New("should init error about existent team"))
	}
}
