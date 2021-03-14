package bind

import (
	"github.com/google/wire"

	"factory/exam/repo"
	"factory/exam/repo/mysql"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	mysql.NewProductMySQLRepo,
	wire.Bind(new(repo.ProductRepoInterface), new(*mysql.ProductMySQLRepo)),
)