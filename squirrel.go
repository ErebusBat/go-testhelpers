package testhelpers

import (
	"fmt"
	"strings"

	sq "github.com/lann/squirrel"
)

const (
	// ES_EXPECTED  = "\nexpected:"
	SQL_FROM     = "\n          FROM "
	SQL_WHERE    = "\n          WHERE "
	SQL_GROUP_BY = "\n          GROUP BY "
	SQL_ORDER_BY = "\n          ORDER BY "
)

// Takes a SQL string and arguments and prints them with ES_SQL and ES_ARGS
func SqlArgsMsg(sql string, args []interface{}) string {
	return fmt.Sprintln(
		ES_SQL, sql,
		ES_ARGS, args)
}

func SqlPrettyArgsMsg(sql string, args []interface{}) string {
	sql = strings.Replace(sql, " FROM ", SQL_FROM, -1)
	sql = strings.Replace(sql, " WHERE ", SQL_WHERE, -1)
	sql = strings.Replace(sql, " GROUP BY ", SQL_GROUP_BY, -1)
	sql = strings.Replace(sql, " ORDER BY ", SQL_ORDER_BY, -1)
	return SqlArgsMsg(sql, args)
}

func SqSelectBuilderAndArgsMsg(sb sq.SelectBuilder) string {
	sql, args, _ := sb.ToSql()
	return SqlPrettyArgsMsg(sql, args)
}
