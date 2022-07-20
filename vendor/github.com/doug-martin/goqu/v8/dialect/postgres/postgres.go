package postgres

import (
	"github.com/doug-martin/goqu/v8"
)

func DialectOptions() *goqu.SQLDialectOptions {
	do := goqu.DefaultDialectOptions()
	do.PlaceHolderRune = '$'
	do.IncludePlaceholderNum = true
	return do
}

func init() {
	goqu.RegisterDialect("postgres", DialectOptions())
}
