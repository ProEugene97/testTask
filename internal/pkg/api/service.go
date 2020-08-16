package api

import (
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/proto"
	"time"
)

type SubService struct {
	db database.IDatabase
}

func NewSubService(db database.IDatabase) *SubService {
	return &SubService {
		db,
	}
}

func (s *SubService) SubscribeOnSportsLines(stream proto.SubService_SubscribeOnSportsLinesServer) error {
	in, err := stream.Recv()
	if err != nil {
		return err
	}

	for {
		ch := make(chan interface{}, 1)
		go func(sports []string, timeOut int32, ch chan interface{}) {
			for {
				select {
				case <-ch:
					return
				default:
				}

				lines, err := s.db.Get(sports)
				if err != nil {
					return
				}


				respLines := make([]*proto.Line, len(lines))
				for i, line := range lines {
					respLines[i] = &proto.Line{
						Sport: line.Sport,
						Coef:  line.Coef,
					}
				}

				err = stream.Send(&proto.SubResponse{Lines: respLines})
				if err != nil {
					return
				}

				timer := time.NewTimer(time.Duration(timeOut) * time.Second)
				<-timer.C
			}
		}(in.Sports, in.Seconds, ch)

		in, err = stream.Recv()
		ch<- 1
		if err != nil {
			return err
		}
	}
}
