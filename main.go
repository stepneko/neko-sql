package main

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stepneko/neko-sql/parser"
)

type Ml struct {
	*parser.BaseMySqlParserListener
	tableNames map[string]struct{}
}

func (m *Ml) EnterTableName(ctx *parser.TableNameContext) {
	if m.tableNames == nil {
		m.tableNames = make(map[string]struct{})
	}
	m.tableNames[ctx.GetText()] = struct{}{}
}
func (m *Ml) GetTableNames() []string {
	arr := make([]string, 0)
	if m.tableNames != nil {
		for k := range m.tableNames {
			arr = append(arr, k)
		}
	}
	return arr
}
func GetTableNames(sql string, sqlType string) []string {
	tokenStream := antlr.NewCommonTokenStream(parser.NewMySqlLexer(antlr.NewInputStream(strings.ToUpper(sql))), antlr.TokenDefaultChannel)
	sqlParser := parser.NewMySqlParser(tokenStream)
	ml := Ml{}
	antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.SqlStatements())
	// switch sqlType {
	// case "dml":
	// 	antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.DmlStatement())
	// case "ddl":
	// 	antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.DdlStatement())
	// }
	return ml.GetTableNames()
}

func main() {
	sql := "create table newtable select * from table1"
	for _, name := range GetTableNames(sql, "ddl") {
		println(name)
	}
}
