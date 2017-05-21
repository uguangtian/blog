package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	View            int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}
type User struct {
	Id   int64
	Name string
}

func init() {
	orm.RegisterModel(new(Category), new(User))
	// orm.RegisterDriver("mysql", orm.DRMMySql)
	orm.RegisterDataBase("default", "mysql", "root:mysql@/blog?charset=utf8")

	orm.RunSyncdb("default", false, true)
}

func CheckConnect() bool {
	o := orm.NewOrm()
	user := User{Name: "Leo"}
	_, err := o.Insert(&user)
	if err != nil {
		log.Fatalf("Insert Error %s\n", err)
	}
	// db, err := sql.Open("mysql", "root:mysql@/blog")
	// if err != nil {
	// 	log.Fatalf("Open database error:%s\n", err)
	// }
	// defer db.Close()
	// query, err := db.Prepare("create table test3(name char(20))")
	// if err != nil {
	// 	log.Fatalf("Prepare sql query error:%s\n", err)
	// }
	// _, err = query.Exec()
	// if err != nil {
	// 	log.Fatalf("Execute sql query error:%s\n", err)
	// }
	//
	return true

}
