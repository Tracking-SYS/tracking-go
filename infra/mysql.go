package infra

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"github.com/Tracking-SYS/tracking-go/utils/envparser"
	"github.com/Tracking-SYS/tracking-go/utils/logger"

	mysqlDriver "github.com/go-sql-driver/mysql"
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
	var cfg mysqlDriver.Config
	var mySQLDSN string
	isTLS := envparser.GetBool("MYSQL_USE_SSL", false)
	if isTLS == true {
		rootCertPool := x509.NewCertPool()
		pem, err := ioutil.ReadFile("./infra/ca.pem")
		if err != nil {
			log.Fatal(err)
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			log.Fatal("Failed to append PEM.")
		}
		mysqlDriver.RegisterTLSConfig("custom", &tls.Config{
			RootCAs: rootCertPool,
		})

		// try to connect to mysql database.
		cfg = mysqlDriver.Config{
			User:                 "sgroot",
			Passwd:               "W-tAdT2k9xc1p7xL",
			Addr:                 "SG-ts-4920-mysql-master.servers.mongodirector.com", //IP:PORT
			Net:                  "tcp",
			DBName:               "database",
			Loc:                  time.Local,
			AllowNativePasswords: true,
			Params: map[string]string{
				"useSSL":                  "true",
				"verifyServerCertificate": "true",
			},
		}
		cfg.TLSConfig = "custom"
		mySQLDSN = cfg.FormatDSN()
	} else {
		mySQLDSN = envparser.GetString("MYSQL_ADDR", "root:123@tcp(localhost:3306)/tracking?charset=utf8&parseTime=True&loc=Local&multiStatements=true")
	}

	return Configuration{
		addr:            mySQLDSN,
		maxOpenConns:    envparser.GetInt("POOL_SIZE", 32),
		maxIdleConns:    envparser.GetInt("MAX_IDLE", 32),
		connMaxLifetime: time.Duration(envparser.GetInt("MAX_LIFETIME", 30)) * time.Minute,
	}
}

//GetAddr ...
func (c Configuration) GetAddr() string {
	return c.addr
}
