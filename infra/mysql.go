package infra

import (
	"time"

	"github.com/Tracking-SYS/tracking-go/utils/envparser"
	"github.com/Tracking-SYS/tracking-go/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Configuration ...
type Configuration struct {
	addr            string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration //minutes
}

//ConnPool ...
type ConnPool struct {
	Conn   *gorm.DB
	config Configuration
}

//GetConnectionPool ...
func GetConnectionPool(config Configuration) (*ConnPool, error) {
	logger := logger.GetLoggerFactory("mysql")
	db, err := gorm.Open(mysql.Open(config.addr), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		logger.Error(err, "Open Mysql Connection was failed")
		return nil, err
	}

	pool, err := db.DB()
	if err != nil {
		logger.Error(err, "Initializing Mysql connection pool")
		return nil, err
	}

	pool.SetMaxOpenConns(config.maxOpenConns)
	pool.SetMaxIdleConns(config.maxIdleConns)
	pool.SetConnMaxLifetime(config.connMaxLifetime)

	return &ConnPool{Conn: db, config: config}, nil
}

//InitConfiguration ...
func InitConfiguration() Configuration {
	return Configuration{
		addr:            envparser.GetString("MYSQL_ADDR", "root:123@tcp(localhost:3306)/tracking?charset=utf8&parseTime=True&loc=Local&multiStatements=true"),
		maxOpenConns:    envparser.GetInt("POOL_SIZE", 32),
		maxIdleConns:    envparser.GetInt("MAX_IDLE", 32),
		connMaxLifetime: time.Duration(envparser.GetInt("MAX_LIFETIME", 30)) * time.Minute,
	}
}

//GetAddr ...
func (c Configuration) GetAddr() string {
	return c.addr
}
