package main

import (
  _ "quickstart/routers"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/mattn/go-sqlite3"
  _ "quickstart/models"
  "fmt"
)


func init() {
    orm.RegisterDriver("sqlite3", orm.DRMySQL)
    err:=orm.RegisterDataBase("default", "sqlite3", "file:data.db")
    if err!=nil{
       panic(err)
    }
}

func main() {
	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := false
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}
        orm.Debug = true
        orm.RunCommand()
	beego.Run()
}

