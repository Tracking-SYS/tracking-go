package bind

import (
	"github.com/google/wire"

	"factory/exam/repo"
	"factory/exam/repo/cache"
	"factory/exam/repo/mysql"
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
