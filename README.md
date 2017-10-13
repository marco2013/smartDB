 smartDbTest.go

package main

import (
	"fmt"
	"smartDB"
)

func main() {

  db, err := smartDB.NewDb("mysql", "root:root@tcp(192.168.1.xx:3306)/test?charset=utf8")

	//query some, return sql.Rows
	rows, _ := db.Query("select id,name from person")

	//query all
	rows2, err := db.Query("select * from person")
    //query with json return
	JsonOut := db.QueryDataRowsToJson("select * from person")
    //insert
	instid, err := db.Insert("insert into person(id,name,age) values(null, 'jack',22)")
    //update
	affNum, err := db.Update("update person set age=100 where name='jack'")
    //delete
	affNum, err = db.Delete("delete from person where name=?", "jack")

}
# smartDB
Encapsulate mysql basic crdu operation in a package with golang

1 need to download golang mysql driver package
go get github.com/go-sql-driver/mysql

