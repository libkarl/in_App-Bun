package app

import (
	"time"
	"fmt"
	"regexp"
	"strings"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/urfave/cli/v2"
)

var RunCMD = &cli.Command{
	Name:  "run",
	Usage: "will try to train the agent",
	Flags: []cli.Flag{
		&cli.StringFlag{ Name: "db", Value: "postgres://charles@localhost:5432/cfapp?sslmode=disable", Usage: "db" },	
	},
	Action: printTest,
}

type User struct {
	name string
	address string
	age int
}

type Hook struct {
	item User

}

type Log struct {
	strip int
}



var spaceReg = regexp.MustCompile(`\s+`)
const startAt = "start_at"

// pokud nejsou data nil udělá v eventu mapu interface -> interface 
func mkEventData(event *bun.QueryEvent) {
	if event.Stash != nil {
		return
	}
	event.Stash = make(map[interface{}]interface{})
}

func (log *Log) truncate(query string) string {
	log.strip = 0
	query = spaceReg.ReplaceAllString(query, " ")
	query = strings.TrimSpace(query)
	if log.strip == 0 {
		return query
	}
	if len(query) <= log.strip {
		return query
	}
	return fmt.Sprintf("%s ....", query[0:log.strip-1])
}


func (h *Log) BeforeQuery(ctx context.Context, e *bun.QueryEvent) context.Context {
	// bofore
	logrus.Debug(e.Query)
	return ctx
}

func (h *Log) AfterQuery(ctx context.Context, e *bun.QueryEvent) {
	query, err := e.QueryAppender.AppendQuery(e.DB.Formatter(), nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	query_string := h.truncate(string(query))
	mkEventData(e)w
	v := e.Stash[startAt].(time.Time)
	diff := time.Since(v)
	lgr := logger.Ctx(e.Ctx).WithField("execution", diff.Seconds())
}

func printTest(ctx *cli.Context) error {

	logrus.SetLevel(logrus.DebugLevel)

	dbstr := ctx.String("db")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbstr)))

	db := bun.NewDB(sqldb, pgdialect.New())

	var hook bun.QueryHook
	hook = &Log{
		strip: 0,
	}
	db.AddQueryHook(hook)


	result, err := db.Exec("select 'healthcheck'")
	if err != nil {
		logrus.Fatal(err)
	}

	//pretty.Println(result)
	logrus.Info(result)


	return nil
}
