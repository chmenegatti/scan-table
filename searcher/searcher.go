package searcher

import (
	"database/sql"
	"fmt"
	"strings"
)

type TableSearcher interface {
	GetTables() ([]string, error)
	SearchTables([]string) ([]string, error)
	SearchTablesWithFields([]string, []string) ([]string, error)
}

type MySQLTableSearcher struct {
	db *sql.DB
}

func (s MySQLTableSearcher) GetTables() ([]string, error) {
	rows, err := s.db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (s MySQLTableSearcher) SearchTables(tables []string) ([]string, error) {
	return nil, fmt.Errorf("Method not implemented")
}

func (s MySQLTableSearcher) SearchTablesWithFields(tables []string, fields []string) ([]string, error) {
	var resultTables []string

	for _, table := range tables {
		query := fmt.Sprintf("SHOW COLUMNS FROM %s", table)
		rows, err := s.db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var found bool
		for rows.Next() {
			var field, _type, null, key string
			var _default, extra sql.NullString
			if err := rows.Scan(&field, &_type, &null, &key, &_default, &extra); err != nil {
				return nil, err
			}
			for _, f := range fields {
				if strings.Contains(field, f) {
					found = true
					break
				}
			}
			if found {
				resultTables = append(resultTables, table)
				break
			}
		}
	}
	return resultTables, nil
}

type TableSearcherFactory struct{}

func (f TableSearcherFactory) CreateSearcher(db *sql.DB) TableSearcher {
	return MySQLTableSearcher{db: db}
}
