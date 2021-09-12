package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Tracking-SYS/tracking-go/utils/envparser"
	"github.com/go-redis/redis/v8"
)

const (
	//FullInfoTTL ...
	FullInfoTTL = 3 * 24 * time.Hour

	//CachePrefix ...
	CachePrefix = "cache"
)

//CacheInteface ...
type CacheInteface interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}) error
}

var redisClient *redis.Client
var _ CacheInteface = &RedisCache{}

//RedisCache ...
type RedisCache struct {
	rdb *redis.Client
}

//NewRedisCacheRepo ...
func NewRedisCacheRepo() *RedisCache {
	db, err := strconv.Atoi(envparser.GetString("REDIS_DB", "0"))
	if err != nil {
		fmt.Printf("read redis db config error: %v\n", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Network:  envparser.GetString("REDIS_NETWORK", "tcp"),
		Addr:     envparser.GetString("REDIS_ADDR", "localhost:6379"),
		Username: envparser.GetString("REDIS_USERNAME", ""),
		Password: envparser.GetString("REDIS_PASS", ""),
		DB:       db,
	})

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
