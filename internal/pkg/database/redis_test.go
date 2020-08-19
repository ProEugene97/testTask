package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"testTask/internal/pkg/models"
	"testing"
)

type TestCaseGet struct {
	Sports []string
	Coefs  []string
	Error  bool
}

func TestHandler_Get(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	rdb := NewRedisDB(pool)

	cases := []TestCaseGet{
		TestCaseGet{
			Sports: []string{"baseball", "soccer"},
			Coefs:  nil,
			Error:  true,
		},
		TestCaseGet{
			Sports: []string{"baseball", "soccer"},
			Coefs:  []string{"1.05", "2.18"},
			Error:  false,
		},
		TestCaseGet{
			Sports: []string{"baseball", "soccer"},
			Coefs:  []string{"2.18"},
			Error:  false,
		},
	}

	for caseNum, item := range cases {
		tmp := make([]interface{}, len(item.Sports))
		for i, sport := range item.Sports {
			tmp[i] = sport
		}

		//var cmd *redigomock.Cmd

		if item.Error {
			conn.Command("MGET", tmp...).ExpectError(fmt.Errorf("error"))
		} else {
			answer := make([]interface{}, len(item.Coefs))
			for i, coef := range item.Coefs {
				answer[i] = coef
			}
			conn.Command("MGET", tmp...).Expect(answer)
		}

		lines, err := rdb.Get(item.Sports)
		if err != nil && !item.Error {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err.Error(), nil)
		}

		if !item.Error {
			for i, line := range lines {
				if line.Sport != item.Sports[i] || line.Coef != item.Coefs[i] {
					t.Errorf("[%d] wrong Line: got %+v, expected %+v",
						caseNum, line, models.Line{Sport: item.Sports[i], Coef: item.Coefs[i]})
				}
			}
		}

	}
}

type TestCaseSet struct {
	Line  models.Line
	Error bool
}

func TestHandler_Set(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	rdb := NewRedisDB(pool)

	cases := []TestCaseSet{
		TestCaseSet{
			Line:  models.Line{Sport: "soccer", Coef: "1.09"},
			Error: true,
		},
		TestCaseSet{
			Line:  models.Line{Sport: "soccer", Coef: "1.09"},
			Error: false,
		},
	}

	for caseNum, item := range cases {

		if item.Error {
			conn.Command("SET", item.Line.Sport, item.Line.Coef).ExpectError(fmt.Errorf("error"))
		} else {
			conn.Command("SET", item.Line.Sport, item.Line.Coef).ExpectError(nil)
		}

		err := rdb.Set(&item.Line)
		if err != nil && !item.Error {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err.Error(), nil)
		}
	}
}

type TestCasePing struct {
	Error bool
}

func TestHandler_Ping(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	rdb := NewRedisDB(pool)

	cases := []TestCasePing{
		TestCasePing{
			Error: true,
		},
		TestCasePing{
			Error: false,
		},
	}

	for caseNum, item := range cases {

		if item.Error {
			conn.Command("PING").ExpectError(fmt.Errorf("error"))
		} else {
			conn.Command("PING").Expect("PONG")
		}

		err := rdb.Ping()
		if err != nil && !item.Error {
			fmt.Println(err)
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err.Error(), nil)
		}
	}
}
