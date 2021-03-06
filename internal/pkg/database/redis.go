package database

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"testTask/internal/pkg/models"
	"time"
)

type RedisDB struct {
	pool *redis.Pool
}

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     80,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewRedisDB(pool *redis.Pool) *RedisDB {
	return &RedisDB{
		pool,
	}
}

func (rd *RedisDB) Get(sports []string) ([]*models.Line, error) {
	conn := rd.pool.Get()
	defer conn.Close()

	tmp := make([]interface{}, len(sports))
	for i, sport := range sports {
		tmp[i] = sport
	}

	coefs, err := redis.Strings(conn.Do("MGET", tmp...))
	if err != nil {
		return nil, errors.Wrap(err, "function Get()")
	}

	l := len(sports)
	if len(sports) > len(coefs) {
		l = len(coefs)
	}

	lines := make([]*models.Line, l)
	for i := 0; i < l; i++ {
		lines[i] = &models.Line{
			Sport: sports[i],
			Coef:  coefs[i],
		}
	}

	return lines, nil
}

func (rd *RedisDB) Set(line *models.Line) error {
	conn := rd.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", line.Sport, line.Coef)
	if err != nil {
		return errors.Wrap(err, "redis error")
	}

	return nil
}

func (rd *RedisDB) Ping() error {
	conn := rd.pool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return errors.Wrap(err, "redis error")
	}

	return nil
}
