package parser

import "github.com/datakit-dev/go-sql-parser/parser/types"

type Database string

const (
	DatabaseAthena      Database = "athena"
	DatabaseBigQuery    Database = "bigquery"
	DatabaseDB2         Database = "db2"
	DatabaseFlinkSQL    Database = "flinksql"
	DatabaseHive        Database = "hive"
	DatabaseMariaDB     Database = "mariadb"
	DatabaseMySQL       Database = "mysql"
	DatabaseNoql        Database = "noql"
	DatabasePostgreSQL  Database = "postgresql"
	DatabaseRedshift    Database = "redshift"
	DatabaseSnowflake   Database = "snowflake"
	DatabaseSqlite      Database = "sqlite"
	DatabaseTransactSQL Database = "transactsql"
	DatabaseTrino       Database = "trino"
)

var Databases = []Database{
	DatabaseAthena,
	DatabaseBigQuery,
	DatabaseDB2,
	DatabaseFlinkSQL,
	DatabaseHive,
	DatabaseMariaDB,
	DatabaseMySQL,
	DatabaseNoql,
	DatabasePostgreSQL,
	DatabaseRedshift,
	DatabaseSnowflake,
	DatabaseSqlite,
	DatabaseTransactSQL,
	DatabaseTrino,
}

func (d Database) String() string {
	return string(d)
}

func (d Database) Supports(s types.Statement) bool {
	// switch s {
	// case types.StatementAlter:
	// 	switch d {
	// 	case DatabaseBigQuery:
	// 		return false
	// 	}
	// case types.StatementCreate:
	// 	switch d {
	// 	case DatabaseBigQuery, DatabaseDB2, DatabaseMariaDB:
	// 		return false
	// 	}
	// case types.StatementDrop:
	// 	switch d {
	// 	case DatabaseBigQuery:
	// 		return false
	// 	}
	// }
	return true
}
