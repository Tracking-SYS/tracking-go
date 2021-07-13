package cache

import (
	"context"
	"fmt"
	"time"

	"encoding/json"

	"github.com/go-redis/redis/v8"
)

const (
	FullInfoTTL = 3 * 24 * time.Hour
	CachePrefix = "cache"
)

//CacheInteface ...
type CacheInteface interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}) error
}

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

var _ CacheInteface = &RedisCache{}

type RedisCache struct {
	rdb *redis.Client
}

//NewRedisCacheRepo ...
func NewRedisCacheRepo() *RedisCache {
	return &RedisCache{
		rdb: redisClient,
	}
}

//Get ...
func (r *RedisCache) Get(ctx context.Context, key string) (model interface{}, err error) {
	pipeline := r.rdb.Pipeline()
	cmd := pipeline.Get(ctx, fmt.Sprintf("%v_%v", CachePrefix, key))
	_, err = pipeline.Exec(ctx)
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, fmt.Errorf("getCache %w", err)
	}

	buf, err := cmd.Bytes()
	if err != nil {
		return nil, fmt.Errorf("read bytes has error: %v", err)
	}

	if buf == nil {
		return nil, nil
	}

	if err = json.Unmarshal(buf, &model); err != nil {
		return nil, fmt.Errorf("unmarshal product: %v", err)
	}

	return model, nil
}

//Set ...
func (r *RedisCache) Set(ctx context.Context, key string, value interface{}) error {
	pipeline := r.rdb.Pipeline()
	cmds := []*redis.StatusCmd{}
	cacheData, err := json.Marshal(&value)
	if err != nil {
		return fmt.Errorf("set cache marshal error: %v", err)
	}

	cmds = append(cmds, pipeline.Set(ctx, fmt.Sprintf("%v_%v", CachePrefix, key), string(cacheData), FullInfoTTL))
	_, err = pipeline.Exec(ctx)
	if err != nil {
		return fmt.Errorf("set cache exec error: %v", err)
	}

	for _, cmd := range cmds {
		if cmd.Err() != nil {
			return fmt.Errorf("set product cache error: %v", cmd.Err())
		}
	}

	return nil
}
