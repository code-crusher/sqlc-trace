package main

import (
	"fmt"
	"strings"

	"github.com/sqlc-dev/sqlc/internal/engine/dolphin"
	"github.com/sqlc-dev/sqlc/internal/engine/postgresql"
	"github.com/sqlc-dev/sqlc/internal/sql/ast"
)

func main() {
	for _, q := range []string{
		"SELECT -a FROM t WHERE a > -1",
		"SELECT a FROM t WHERE NOT a",
	} {
		stmts, err := postgresql.NewParser().Parse(strings.NewReader(q + ";"))
		if err != nil {
			fmt.Println("ERR", err)
			continue
		}
		fmt.Printf("pg %s => %s\n", q, ast.Format(stmts[0].Raw, postgresql.NewParser()))
	}
	for _, q := range []string{
		"SELECT id FROM t WHERE id > -5 AND id > ~1 AND NOT active",
	} {
		stmts, err := dolphin.NewParser().Parse(strings.NewReader(q + ";"))
		if err != nil {
			fmt.Println("ERR", err)
			continue
		}
		fmt.Printf("my %s => %s\n", q, ast.Format(stmts[0].Raw, dolphin.NewParser()))
	}
}
