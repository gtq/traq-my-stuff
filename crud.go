package main

import (
	_ "fmt"
	_ "reflect"
	_ "net/http"
        "database/sql"
        _ "github.com/mattn/go-sqlite3"
)
func OpenDB() *sql.DB {
        db, err := sql.Open("sqlite3", "./traqstuff")
        checkErr(err)
	return db
}
func CreateEntry(querystr string, fields []string, datamap map[string]interface{}) int64 {
vals := make([]interface{}, len(fields))
for key, _ := range fields {
	vals[key] = datamap[fields[key]]
}
//for key := range fields {
//	fmt.Println(key,datamap[fields[key]],reflect.TypeOf(datamap[fields[key]]))
//}
        db := OpenDB()
        stmt, err := db.Prepare(querystr)
        checkErr(err)

        res, err := stmt.Exec(vals... )
        checkErr(err)

	id, err := res.LastInsertId()
	//fmt.Fprintln(len(datamap))
	//for item := range len(datamap) {
	//	fmt.Fprintln(datamap[item])
	//}
	//fmt.Fprintln(datamap["name"])
	return id
}

func ReadEntry(querystr string,id int) []map[string]interface{} {

        db := OpenDB()
        stmt, err := db.Prepare(querystr)
        checkErr(err)
        rows, err := stmt.Query(id)
        checkErr(err)
        defer rows.Close()

	return IterMapRows(rows)

}
func IterMapRows(rows *sql.Rows) []map[string]interface{} {
        columns, err := rows.Columns()
        checkErr(err)
        tableData := make([]map[string]interface{}, 0)
        count := len(columns)
        values := make([]interface{}, count)
        scanArgs := make([]interface{}, count)
        for i := range values {
                scanArgs[i] = &values[i]
        }

        for rows.Next() {
                err := rows.Scan(scanArgs...)
                checkErr(err)

                entry := make(map[string]interface{})
                for i, col := range columns {
                        v := values[i]

                        b, ok := v.([]byte)
                        if (ok) {
                                entry[col] = string(b)
                        } else {
                                entry[col] = v
                        }
                }

                tableData = append(tableData, entry)
        }

	return tableData

}

func UpdateEntry(querystr string, id int, fields []string, datamap map[string]interface{}) sql.Result {
vals := make([]interface{}, len(fields))
for key, _ := range fields {
        vals[key] = datamap[fields[key]]
	if (fields[key] == "id") {
		vals[key] = id
	}
//	fmt.Println(key,fields[key])
}

        db := OpenDB()
        stmt, err := db.Prepare(querystr)
        checkErr(err)

        res, err := stmt.Exec(vals... )
        checkErr(err)
//return 1
	return res
}

func DeleteEntry(querystr string, id int) sql.Result {
	db := OpenDB()
	stmt, err := db.Prepare(querystr)
        checkErr(err)

        res, err := stmt.Exec(id)
        checkErr(err)
	//fmt.Fprintln(w,reflect.TypeOf(res))
	//return 1
	return res 
}


func EntryList(querystr string) []map[string]interface{} {
        db := OpenDB()
        rows, err := db.Query(querystr)
        checkErr(err)

	return IterMapRows(rows)
}
