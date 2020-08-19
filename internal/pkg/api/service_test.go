package api

import (
	"context"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	dbMock "testTask/internal/pkg/database/mock"
	"testTask/internal/pkg/logger"
	"testTask/internal/pkg/models"
	"testTask/internal/pkg/proto"
	"testing"
)

type TestCaseSub struct {
	AmountOfSending   int
	AmountOfReceiving int
	SendData          []*proto.SubRequest
	GetData           []*models.Line
	Coefs             [][]string
	Error             bool
}

func TestHandler_SubscribeOnSportsLines(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := logger.NewLogger("DEBUG")
	db := dbMock.NewMockIDatabase(ctrl)

	lis := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	proto.RegisterSubServiceServer(server, NewSubService(db, l))

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	dialFinc := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialFinc), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	defer cancel()

	client := proto.NewSubServiceClient(conn)
	stream, err := client.SubscribeOnSportsLines(ctx)
	if err != nil {
		t.Fatalf("No connection: %v", err)
	}

	cases := []TestCaseSub{
		{
			AmountOfSending:   2,
			AmountOfReceiving: 2,
			SendData: []*proto.SubRequest{
				&proto.SubRequest{
					Sports:  []string{"soccer", "football"},
					Seconds: 1,
				},
				&proto.SubRequest{
					Sports:  []string{"baseball", ""},
					Seconds: 2,
				},
			},
			Coefs: [][]string{{"1.53", "2.03"}, {"1.85", "1.93"}},
			Error: true,
		},
	}

	for caseNum, item := range cases {
		for _, req := range item.SendData {
			for _, coef := range item.Coefs {
				answer := make([]*models.Line, len(req.Sports))
				for k := range answer {
					answer[k] = &models.Line{Sport: req.Sports[k], Coef: coef[k]}
				}
				gomock.InOrder(db.EXPECT().Get(req.Sports).Return(answer, nil))
			}
		}
		for _, req := range item.SendData {
			err := stream.Send(req)
			if err != nil {
				t.Errorf("[%d] wrong Error: got %+v, expected %+v",
					caseNum, err, nil)
			}

			for _, coef := range item.Coefs {
				resp, err := stream.Recv()
				if err != nil {
					t.Errorf("[%d] wrong Error: got %+v, expected %+v",
						caseNum, err, nil)
				}

				for i, line := range resp.Lines {
					if line.Sport != req.Sports[i] || line.Coef != coef[i] {
						t.Errorf("[%d] wrong Response: got %+v, expected %+v",
							caseNum, line, proto.Line{Sport: req.Sports[i], Coef: coef[i]})
					}
				}
			}
		}

		cancel()
		_, err := stream.Recv()
		if err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, nil, "stream canceled")
		}

	}
}
