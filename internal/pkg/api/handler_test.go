package api

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"strings"
	"testTask/internal/pkg/database/mock"
	"testTask/internal/pkg/logger"
	"testTask/internal/pkg/models"
	"testing"
)

type TestCaseStatus struct {
	IsReady   bool
	Counter   int
	Workers   int
	MockError bool
	Status    int
	Response  string
}

func TestHandler_Status(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	dbMock := mock.NewMockIDatabase(ctl)

	l := logger.NewLogger("DEBUG")

	cases := []TestCaseStatus{
		TestCaseStatus{
			IsReady:   false,
			Counter:   0,
			Workers:   3,
			MockError: true,
			Status:    202,
			Response: `{"status":"storages are not synchronized"}
`,
		},
		TestCaseStatus{
			IsReady:   false,
			Counter:   3,
			Workers:   3,
			MockError: true,
			Status:    202,
			Response: `{"status":"database is not available"}
`,
		},
		TestCaseStatus{
			IsReady:   true,
			Counter:   3,
			Workers:   3,
			MockError: false,
			Status:    200,
			Response: `{"status":"OK"}
`,
		},
	}

	for caseNum, item := range cases {
		handler := NewHandler(dbMock, l, &item.Counter, item.Workers)
		r := httptest.NewRequest("GET", "/ready", strings.NewReader(""))
		w := httptest.NewRecorder()

		ctx := r.Context()
		reqId := fmt.Sprintf("%016x", rand.Int())[:10]
		ctx = context.WithValue(ctx, models.ContextKey{}, reqId)
		r = r.WithContext(ctx)

		if item.Counter >= item.Workers {
			if item.MockError {
				gomock.InOrder(
					dbMock.EXPECT().Ping().Return(fmt.Errorf("error")),
				)
			} else {
				gomock.InOrder(
					dbMock.EXPECT().Ping().Return(nil),
				)
			}
		}

		handler.Status(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.Status != resp.StatusCode {
			t.Errorf("[%d] wrong status Response: got %+v, expected: %v",
				caseNum, resp.Status, item.Status)
		}

		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}
