package bind

import (
	"github.com/google/wire"

	"github.com/Tracking-SYS/tracking-go/repo"
	"github.com/Tracking-SYS/tracking-go/repo/cache"
	"github.com/Tracking-SYS/tracking-go/repo/mysql"
)

//GraphSet Repo
var GraphSet = wire.NewSet(
	mysql.NewProductMySQLRepo,
	wire.Bind(new(repo.ProductRepoInterface), new(*mysql.ProductMySQLRepo)),

	mysql.NewTaskMySQLRepo,
	wire.Bind(new(repo.TaskRepoInterface), new(*mysql.TaskMySQLRepo)),

	cache.NewRedisCacheRepo,
	wire.Bind(new(cache.CacheInteface), new(*cache.RedisCache)),
)
