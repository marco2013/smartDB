// smartDB project smartDB.go
package smartDB

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type MyDb struct {
	*sql.DB
}

func NewDb(driverName, dataSourceName string) (*MyDb, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	mydb := &MyDb{DB: db}
	return mydb, nil
}

func (this *MyDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r, err := this.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return r, nil
}

//queryRows2Json
func (this *MyDb) QueryDataRowsToJson(query string, args ...interface{}) string {
	rows, err := this.Query(query, args...)
	if err != nil {
		return ""
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		//fmt.Println("row to json:", err)
		return ""
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		//fmt.Println("2 row to json:", err)
		return ""
	}

	//fmt.Println(string(jsonData))
	return string(jsonData)
}

/*format:  insert into test(id,name) values(null, '444') */
func (this *MyDb) Insert(strSql string) (int64, error) {
	result, err := this.Exec(strSql)
	if err != nil {
		return 0, err
	}
	ins_id, _ := result.LastInsertId()
	return ins_id, nil
}

/*根据主键更新数据记录,返回所影响的行数
("update test set name = '000' where id > ?", 2); */
func (this *MyDb) Update(strSql string, args ...interface{}) (int64, error) {

	result, err := this.Exec(strSql, args...)
	if err != nil {
		return 0, err
	}
	var affNum int64
	affNum, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affNum, nil
}

/*根据主键删除一条数据,返回所影响的行
("delete from test where id = ?", 3); */
func (this *MyDb) Delete(strSql string, args ...interface{}) (int64, error) {
	result, err := this.Exec(strSql, args...)
	if err != nil {
		return 0, err
	}
	var affNum int64
	affNum, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affNum, nil
}
