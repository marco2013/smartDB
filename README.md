//first
need to download golang mysql driver package and smartDB package
go get github.com/go-sql-driver/mysql
go get github.com/marco2013/smartDB
//second use this package in your projects

smartDbTest.go

package main

import (
	"fmt"
	"github.com/marco2013/smartDB"
)

func main() {

  db, err := smartDB.NewDb("mysql", "usrname:pwd@tcp(xx.xx.xx.xx:3306)/test?charset=utf8")

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


