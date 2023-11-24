package caching

import (
	"context"
	"encoding/json"
	"errors"
	"job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)
//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=caching
type Redis struct {
	rdb *redis.Client
}
type Cache interface {
	AddCache(ctx context.Context, jobid uint, jobData models.Job) error
	GetCache(ctx context.Context, jobid uint) (string, error)
}

func NewRedis(rdb *redis.Client) (Cache, error) {
	if rdb == nil {
		log.Info().Msg("redis cannot be nil")
		return nil, errors.New("---------------")
	}
	return &Redis{
		rdb: rdb,
	}, nil
}
func (re *Redis) AddCache(ctx context.Context, jobid uint, jobData models.Job) error {
	jobID := strconv.FormatUint(uint64(jobid), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}
	err = re.rdb.Set(ctx, jobID, val, 15*time.Minute).Err()

	return err
}
func (re *Redis) GetCache(ctx context.Context, jobid uint) (string, error) {
	jobID := strconv.FormatUint(uint64(jobid), 10)
	str, err := re.rdb.Get(ctx, jobID).Result()
	return str, err
}
