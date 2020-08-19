package api

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"io"
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/proto"
	"time"
)

type SubService struct {
	db     database.IDatabase
	logger *zap.Logger
}

func NewSubService(db database.IDatabase, logger *zap.Logger) *SubService {
	return &SubService{
		db,
		logger,
	}
}

func (s *SubService) SubscribeOnSportsLines(stream proto.SubService_SubscribeOnSportsLinesServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	eg, ctx := errgroup.WithContext(ctx)

	chData := make(chan *proto.SubRequest, 1)

	eg.Go(func() error {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return nil
			default:
			}
			in, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				return nil
			default:
				chData <- in
			}
		}
	})

	eg.Go(func() error {
		defer cancel()

		sports := []string{}
		var timeout int32 = 0
		for {
			time.Sleep(time.Duration(timeout) * time.Second)

			select {
			case sub := <-chData:
				timeout = sub.Seconds
				sports = sub.Sports
			case <-ctx.Done():
				return nil
			default:
			}

			if len(sports) == 0 {
				continue
			}

			lines, err := s.db.Get(sports)
			if err != nil {
				s.logger.Error(err.Error(),
					zap.String("func", "db.Get"),
					zap.Strings("sports", sports),
				)
				continue
			}

			respLines := make([]*proto.Line, len(lines))
			for i, line := range lines {
				respLines[i] = &proto.Line{
					Sport: line.Sport,
					Coef:  line.Coef,
				}
			}

			s.logger.Debug("sent Lines",
				zap.String("func", "SubscribeOnSportsLines"),
				zap.Any("lines", respLines),
			)

			err = stream.Send(&proto.SubResponse{Lines: respLines})
			if err != nil {
				return err
			}
		}
	})

	return eg.Wait()
}
