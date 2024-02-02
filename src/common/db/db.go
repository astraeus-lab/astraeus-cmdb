// Package db provides clients that can connect to various types of databases.
//
// More detail: https://gorm.io/docs
package db

import (
	"time"

	"gorm.io/gorm"
)

var defaultDBConnect *gorm.DB

// InitDBConnectPool initialize database connection based on database type.
//
// Turn off automatic table creation in the program,
// table management should rely on external initialization.
func InitDBConnectPool(dbType, host, user, passwd, dbName string, maxConn, maxIdelConn, coonMaxIdel int) (err error) {
	defaultDBConnect, err = initDefaultConnectByDBType(newDBConnectParam(dbType, user, passwd, host, dbName))
	if err != nil {
		return
	}

	connect, err := defaultDBConnect.DB()
	if err != nil {
		return
	}

	connect.SetMaxOpenConns(maxConn)
	connect.SetMaxIdleConns(maxIdelConn)
	connect.SetConnMaxIdleTime(time.Duration(coonMaxIdel) * time.Minute)

	return
}

// GetDBConnect get db connection.
// Created by initializing configuration parameters.
//
// After calling, *sql.DB's Close() function cannot be used
// to close the connection, which will cause the database connection
// pool to close, resulting in no database connection available.
func GetDBConnect() *gorm.DB {

	return defaultDBConnect
}
