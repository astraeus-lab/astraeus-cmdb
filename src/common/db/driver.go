package db

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	MySQLType      = "MySQL"
	PostgreSQLType = "PostgreSQL"

	MySQLDSN      = "%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4"
	PostgreSQLDSN = "user=%s password=%s host=%s port=%s dbname=%s"
)

type dbConenctParam struct {
	dbType  string
	user    string
	passwd  string
	address string
	port    string
	dbName  string
}

func newDBConnectParam(dbType, user, passwd, host, dbName string) *dbConenctParam {
	addressAndPort := strings.Split(host, ":")
	return &dbConenctParam{
		dbType:  dbType,
		user:    user,
		passwd:  passwd,
		address: addressAndPort[0],
		port:    addressAndPort[1],
		dbName:  dbName,
	}
}

// initDefaultConnectByDBType connect different type of databases based on parameter.
func initDefaultConnectByDBType(param *dbConenctParam) (res *gorm.DB, err error) {
	switch strings.ToUpper(param.dbType) {
	case strings.ToUpper(MySQLType):
		return initMySQLConnect(MySQLDSN, param)

	case strings.ToUpper(PostgreSQLType):
		return initPostgreSqlConnect(PostgreSQLDSN, param)

	default:
		return initMySQLConnect(MySQLDSN, param)
	}
}

func initMySQLConnect(dbDSN string, param *dbConenctParam) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf(dbDSN, param.user, param.passwd, param.address, param.port, param.dbName),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), getGORMConnectConfig())
}

func initPostgreSqlConnect(dbDSN string, param *dbConenctParam) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf(dbDSN, param.user, param.passwd, param.address, param.port, param.dbName),
		PreferSimpleProtocol: true,
	}), getGORMConnectConfig())
}

func getGORMConnectConfig() *gorm.Config {
	return &gorm.Config{
		Logger:               logger.Discard,
		NowFunc:              time.Now().Local,
		NamingStrategy:       &schema.NamingStrategy{SingularTable: true},
		AllowGlobalUpdate:    false,
		DisableAutomaticPing: false,
		FullSaveAssociations: true,
	}
}
