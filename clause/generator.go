package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...any) (string []any)

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
}

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

func _insert(values ...any) (string, []any) {
	// INSERT INTO $tableName ($fields)
	tableName := values[0]
	fileds := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fileds), []any{}
}

func _values(values ...any) (string []any) {
	// VALUES ($v1),($v2),...
}

func _select(values ...any) (string []any) {

}

func _limit(values ...any) (string []any) {

}

func _where(values ...any) (string []any) {

}
func _orderBy(values ...any) (string []any) {

}
