package parser

type Database string

const (
	Athena      Database = "athena"
	BigQuery    Database = "bigquery"
	DB2         Database = "db2"
	FlinkSQL    Database = "flinksql"
	Hive        Database = "hive"
	MariaDB     Database = "mariadb"
	MySQL       Database = "mysql"
	Noql        Database = "noql"
	PostgreSQL  Database = "postgresql"
	Redshift    Database = "redshift"
	Snowflake   Database = "snowflake"
	Sqlite      Database = "sqlite"
	TransactSQL Database = "transactsql"
	Trino       Database = "trino"
)

var Databases = []Database{
	Athena,
	BigQuery,
	DB2,
	FlinkSQL,
	Hive,
	MariaDB,
	MySQL,
	Noql,
	PostgreSQL,
	Redshift,
	Snowflake,
	Sqlite,
	TransactSQL,
	Trino,
}

func (d Database) String() string {
	return string(d)
}
