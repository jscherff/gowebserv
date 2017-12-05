// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	`database/sql`
	`fmt`
	`strings`

	`github.com/jscherff/cmdbd/utils`
)

// Queries contains SQL queries, column lists, and prepared statements.
type Queries struct {
	Query map[string][]string
	Cols map[string][]string
	Stmt map[string]*sql.Stmt
}

// NewQueries creates and initializes a new Queries instance.
func NewQueries(cf string, db *Database) (this *Queries, err error) {

	this = &Queries{
		Query: make(map[string][]string),
		Cols: make(map[string][]string),
		Stmt: make(map[string]*sql.Stmt),
	}

	if err = utils.LoadConfig(this, cf); err != nil {
		return nil, err
	}

	for key, query := range this.Query {

		rows, err := db.Query(`CALL proc_usbCi_list_columns(?)`, query[1])

		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {

			var col string

			if err = rows.Scan(&col); err != nil {
				return nil, err
			}

			this.Cols[key] = append(this.Cols[key], col)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		var sql string

		switch query[0] {

		case `INSERT_EMPTY`:

			sql = fmt.Sprintf(`INSERT INTO %s VALUES ()`,
				query[1],
			)

		case `INSERT_ALL`:

			params := strings.Repeat(`?,`, len(this.Cols[key])-1) + `?`

			sql = fmt.Sprintf(`INSERT INTO %s VALUES (%s)`,
				query[1],
				params,
			)

		case `UPDATE_LIST`:

			sql = fmt.Sprintf(`UPDATE %s SET %s WHERE %s`,
				query[1],
				query[2],
				query[3],
			)

		case `SELECT_ALL`:

			sql = fmt.Sprintf(`SELECT * FROM %s WHERE %s`,
				query[1],
				query[2],
			)

		case `SELECT_LIST`:

			sql = fmt.Sprintf(`SELECT %s FROM %s WHERE %s`,
				query[1],
				query[2],
				query[3],
			)

		case `REPLACE_LIST`:

			sql = fmt.Sprintf(`REPLACE INTO %s (%s) VALUES (%s)`,
				query[1],
				query[2],
				query[3],
			)
		}

		if this.Stmt[key], err = db.Prepare(sql); err != nil {
			return nil, err
		}
	}

	return this, nil
}

// Close closes the prepared statements.
func (this *Queries) Close() {
	for _, stmt := range this.Stmt {
		stmt.Close()
	}
}
